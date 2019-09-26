package mdc

import (
  "sync"
  "fmt"
  "github.com/tealeg/xlsx"
  "github.com/iancoleman/orderedmap"
)


//序号	站点	设备	信号名称	信号值	时间
//881	右侧面视图	AMM_PMAC725_1	AB线电压	387.71V	2018-05-01 03:10:32

func Mdc(pwd string, wg *sync.WaitGroup)(){
  fmt.Printf("-----------开始 %s 处理------------\n", pwd)

  //TODO: 这里读取excle表格的数据为60MB往上？ 所以很慢数据量太大了
  excel, err := xlsx.OpenFile(pwd)
  if err!=nil{
    fmt.Printf("open excel error:%s", err.Error())
    return
  }
  var MDCMAP = orderedmap.New()
  for _, sheet := range excel.Sheets{
    for i, _ := range sheet.Rows{
      if i!=0 && sheet.Cell(i, 0).Value != EMPTY{
        deviceName := sheet.Cell(i, 2).Value
        spotName := sheet.Cell(i, 3).Value
        value := sheet.Cell(i, 4).Value
        updateTimes := sheet.Cell(i, 5).Value
        // 按设备区分
        if val, ok := MDCMAP.Get(deviceName); ok{
          TMP := val.(*Infos)
          // 对比测点名称 重复则对比时间 取最早的
          if TMP.SpotName == spotName{
            break
          }
        }
        // 设置新值
        MDCMAP.Set(deviceName, &Infos{
          DeviceName:deviceName,
          SpotName:spotName,
          Value:value,
          UpdateTime:updateTimes,
        })
      }
    }
    return
  }
  fmt.Println(MDCMAP)

  // 通知协程停止
  wg.Done()
  fmt.Printf("-----------处理 %s 完毕------------\n", pwd)
}
