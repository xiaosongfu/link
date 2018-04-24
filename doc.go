/*
 link 项目用来收集想要收藏的网页链接

 1. 可以给网页链接加分类
 2. 可以给网页链接加 Tag

//---------------------------
 API 文档

 1. 操作 Category
 1.1 添加 Category
	api：/api/v1/addCategory
	method: post
	params：categoryId categoryName
 1.2 获取所有 Category
	api：/api/v1/getCategorys
	method: get
    params：-

 2. 操作 Link
 2.1 添加 link
	api：/api/v1/addLink
	method: post
    params：url categoryId tag
 备注：如果没有传 categoryId 或为 ""，则使用默认的 categoryId
      tag 可以为 ""
 2.2 获取所有 link
	api：/api/v1/getLinks
	method: get
    params：-
 2.3 根据 categoryId 获取 link
	api：/api/v1/getLinksByCategoryId
	method: get
    params：categoryId
 备注：如果没有传 categoryId 或为 ""，则使用默认的 categoryId
//---------------------------
*/
package main
