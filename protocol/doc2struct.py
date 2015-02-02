#! /usr/bin/python
# encoding: UTF-8

import sys
reload(sys)
sys.setdefaultencoding('utf-8')
import os
import urllib2
import urlparse
import lxml.html
import json
import pprint

addresp={
  "GetServiceCodeList": ["CustomerCode"],
  "Common": ["ErrorType", "ErrorMessage"],
}

def gettbl(root, title):
  paramtitle=root.xpath(u'//div[@class="section"]/h2[text()="%s"]' %(title))
  if len(paramtitle)==0: return None
  params=paramtitle[0].getparent()
  heads=dict(map(lambda f: (f.get("id"), f.text_content().strip()), params.xpath('.//th')))
  names=map(lambda f: f.text_content().strip(), params.xpath('.//th'))
  opt={}
  pent=[]
  for p in params.xpath('.//td'):
    hdrid=p.get("headers").strip()
    if hdrid in heads:
      pent.append(" ".join(map(lambda f: f.strip(), p.text_content().split("\n"))).strip())
      if names[-1]==heads[hdrid] and pent[0]!='':
        opt[pent[0]]=pent[1:]
        pent=[]
  return names, opt

def typefix(s):
  tbl={"String":"string", "Stinrg":"string", "Array":"[]string"}
  if s in tbl:
    return tbl[s]
  return s

def tbl2dict(opt):
  print "tbl2dict from:"
  for k in sorted(opt.keys()):
    print k, "/".join(opt[k])
  res={}
  for k in sorted(opt.keys()):
    v=opt[k]
    kk=k.split(".")
    typ=v[0]
    cur=res
    while len(kk)!=1:
      if kk[0] not in cur or type(cur[kk[0]])!=type({}):
        cur[kk[0]]={}
      cur=cur[kk[0]]
      kk=kk[1:]
    if type(cur)==type([]):
      cur
    cur[kk[0]]=v
  return res

def dump2struct(ofp, tbl):
  ofp.write("{\n")
  for k,v in tbl.items():
    if type(v)==type({}):
      if k.endswith("List"):
        ofp.write("%s []struct" %(k))
      else:
        ofp.write("%s struct" %(k))
      dump2struct(ofp, v)
    elif type(v)==type([]):
      print k, v
      isar=""
      if k.endswith("List") and v[0]!="Array":
        isar="[]"
      ofp.write(" %s %s%s // %s\n" %(k, isar, typefix(v[0]), v[1]))
    elif v==None:
      ofp.write(" %s\n" %(k))
  ofp.write("}\n")

def apiconvert(root, url, p=[u"リクエストパラメータ", u"レスポンス"]):
  metatag=root.xpath('//meta[@name="DC.Title"]')
  if len(metatag)!=1: return None
  name=metatag[0].get("content")
  if name==u"API利用方法":
    name="Common"
  print "name=", name
  ofp=open(name+".go", "w")
  print >>ofp, "// ", url
  print >>ofp, ""
  print >>ofp, "package protocol"
  # request
  names, opt=gettbl(root, p[0])
  for k in opt.keys():
    if opt[k][0]==u"○":
      opt[k][0]=True
    else:
      opt[k][0]=False
  print >>ofp, "type %sArg struct {" %(name)
  for k in sorted(opt.keys()):
    kname=k.rstrip(".#")
    if opt[k][0]:
      tag="`gpapi:\"%s,required\"`" %(kname)
    else:
      tag="`gpapi:\"%s,optional\"`" %(kname)
    kname=kname.replace(".", "_")
    if k.endswith(".#"):
      print >>ofp, " ", kname, "[]string", tag, "//", opt[k][1]
    elif opt[k][1].find(u"整数")!=-1:
      print >>ofp, " ", kname, "int", tag, "//", opt[k][1]
    else:
      print >>ofp, " ", kname, "string", tag, "//", opt[k][1]
  print >>ofp, "}"
  print >>ofp, ""
  # response
  names, opt=gettbl(root, p[1])
  odict=tbl2dict(opt)
  assert len(odict)==1
  firstkey=odict.keys()[0]
  if name in addresp:
    for i in addresp[name]:
       if name==u"Common":
           odict[i]=["string", ""]
       else:
           odict[firstkey][i]=["string", "undocumented"]
  # pprint.pprint(odict)
  if type(odict[firstkey])==type({}):
    ofp.write("type %sResponse struct" %(name))
    odict[firstkey]["CommonResponse"]=None
    dump2struct(ofp, odict[firstkey])
  else:
    # Echo, Common
    ofp.write("type %sResponse struct" %(name))
    if name!=u"Common":
      odict["CommonResponse"]=None
    dump2struct(ofp, odict)

def apiversion(root):
  r=map(lambda f: f.text, root.xpath('//td[@headers="d61843e2373 "]'))
  if len(r)!=0: return r[0]
  return None

def apilist(root):
  paramtitle=root.xpath(u'//div[@class="section"]/h2[text()="API一覧"]')
  if len(paramtitle)==0: return None
  params=paramtitle[0].getparent()
  idn=params.xpath('.//th[text()="Action"]')[0].get("id")
  print "idn", idn
  res={}
  for a in params.xpath('.//td[@headers="%s"]/a' %(idn+" ")):
    res[a.text.strip()]=a.get("href")
  return res

def getfp(baseurl, fname):
  url=urlparse.urljoin(baseurl, fname)
  if not os.path.exists(fname):
    fp=file(fname, "w")
    fp.write(urllib2.urlopen(url).read())
  return file(fname), url

def main(args):
  docbase="http://manual.iij.jp/gp/gpapi/"
  #docbase="file://"+os.getcwd()+"/"
  apicommon_html="9564008.html"
  commonp=[u"APIリクエスト時の共通パラメータ", u"レスポンスパラメータ"]
  if len(args)==0:
    apilist_html="8696725.html"
    fp, url=getfp(docbase, apilist_html)
    root=lxml.html.parse(fp)
    apiver=apiversion(root)
    apis=apilist(root)
    print "\n".join(map(lambda f: urlparse.urljoin(docbase, f), apis.values()))
    # common
    fp, url=getfp(docbase, apicommon_html)
    apiroot=lxml.html.parse(fp)
    apiconvert(apiroot, url, commonp)
    # each
    for api, href in apis.items():
      fp, url=getfp(docbase, href)
      apiroot=lxml.html.parse(fp)
      apiconvert(apiroot, url)
  else:
    for href in args:
      fp, url=getfp(docbase, href)
      apiroot=lxml.html.parse(fp)
      if href==apicommon_html:
        apiconvert(apiroot, url, commonp)
      else:
        apiconvert(apiroot, url)
  os.system("go fmt .")
  os.system("go vet .")
  os.system("go fix .")

if __name__=="__main__":
  main(sys.argv[1:])
