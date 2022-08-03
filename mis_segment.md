
```js

// web local path
var request = $WEB.getParamsTable().getHttpRequest();
println(request.getSession().getServletContext().getRealPath("/"));
String objectName = fullPath.substring(fullPath.indexOf("/uploads")+1);

```