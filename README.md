# rmcomment
删除代码中的注释

使用的有限状态机实现的这个功能，目前的检测还是比较简单，后续再继续加强。

### 使用
InitRm方法传入单行注释标志，块注释开始标志，块注释结束标志
```go
  rmcomment.InitRm("--", "/*", "*/")
	s := `
		/* comment
		adfaf */
		content

	`
	rmcomment.Rm(s)

```
