package main

import (
  "fmt"
  "io/ioutil"
  "flag"
  "strings"
  "os"
  "sync"

  "path/filepath"
  "dome/mdc"
)

var(
  dir *string
)

const(
  MDCSTR = "MDC"
)

func parse(){
  dir = flag.String("d", `F:\GO\src\dome\templates`, "file path.")
  flag.Parse()
}

// 处理mdc表格信息 一个目录开启一个协程
func runMdc(arr []os.FileInfo){
  wg := &sync.WaitGroup{}
  for _, val :=range arr{
    if strings.Contains(strings.ToUpper(val.Name()), MDCSTR){
      pwd := filepath.Join(*dir, val.Name())
      list, err := ioutil.ReadDir(pwd)
      if err != nil{
        fmt.Println(err)
        return
      }
      for _, info := range list{
        filepwd := filepath.Join(pwd, info.Name())
        fileList, err := ioutil.ReadDir(filepwd)
        if err != nil{
          fmt.Println(err)
          return
        }
        for _, data := range fileList{
          wg.Add(1)
          go mdc.Mdc(filepath.Join(filepwd, data.Name()), wg)
          // 先拿一个做实验 完成后直接结束
          goto LOOP
        }
      }
    }
  }
  LOOP:
  wg.Wait()
}

func init(){

}


func main(){
  // parse input param
  parse()

  // get dir list
  arr, err := ioutil.ReadDir(*dir)
  if err != nil{
    fmt.Println(err)
    return
  }

  // start mdc table
  runMdc(arr)
}
