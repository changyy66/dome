package mdc

import (
  "fmt"
  "github.com/iancoleman/orderedmap"
  "github.com/tealeg/xlsx"
  "path/filepath"
  "strings"
)

// 创建新的excel


// 获取新表格名称
func GetNewExcelName(pwd string)string{
  arr := strings.Split(pwd, `\`)
  arr[len(arr)-1] = "new" + arr[len(arr)-1]
  return filepath.Join(arr...)
}

func WriteData(sheet *xlsx.Sheet, MDC *orderedmap.OrderedMap){
  for _, val := range MDC.Keys(){
   data, _ := MDC.Get(val)
   spot := data.(*orderedmap.OrderedMap)
   for _, info := range spot.Keys(){
     fmt.Println(info)
   }
  }
}