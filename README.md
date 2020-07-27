# board-golang

go 언어로 간단한 게시판을 구현해보려 합니다.  
Simple board by golang  

go에서 사용된 라이브러리는 `/go.mod`를 참조해주세요.  
Plz references `/go.mod` for used golang libraries

Vue.js에서 사용된 라이브러리는 `/view/package.json`를 참조해주세요.  
Plz references `/view/package.json` for used vue libraries

### Tech Stack
- [x] golang
- [x] MySql (by docker)
- [x] Vue.js 

### Use
1. `go build main.go`  
2. `npm run serve`
3. `localhost:8080` in browser
4. DONE!  

> golang default port is 8090.  
> It can be set in main.go `os.Setenv("PORT", "8090")`
