package mdc

import (
  "sync"
  "fmt"
  "github.com/tealeg/xlsx"
  "github.com/iancoleman/orderedmap"
  "time"
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
  // 获取sheet表格
  for _, sheet := range excel.Sheets{
    // 获取行数
    for i, _ := range sheet.Rows{
      if i!=0 && sheet.Cell(i, 0).Value != EMPTY{
        deviceName := sheet.Cell(i, 2).Value
        spotName := sheet.Cell(i, 3).Value
        value := sheet.Cell(i, 4).Value
        updateTimes := sheet.Cell(i, 5).Value
        // 按设备区分
        if val, ok := MDCMAP.Get(deviceName); ok{
          TMP := val.(*MDCInfo)
          // 对比测点名称
          if TMP.SpotName == spotName{
            // 重复则对比时间 取最早的
            if diffTime(TMP.UpdateTime, updateTimes){
              break
            }
          }
        }
        // 设置新值
        MDCMAP.Set(deviceName, &MDCInfo{
          DeviceName:deviceName,
          SpotName:spotName,
          Value:value,
          UpdateTime:updateTimes,
        })
      }
    }
  }
  fmt.Println(MDCMAP)

  // 通知协程停止
  wg.Done()
  fmt.Printf("-----------处理 %s 完毕------------\n", pwd)
}


// 把两个字符串转成时间类型做对比
// t1大于t2 返回false 反之true
func diffTime(t1, t2 string)bool{
  loc, _ := time.LoadLocation("Local")
  t1Time, _ := time.ParseInLocation(GOTIME, t1, loc)
  t2Time, _ := time.ParseInLocation(GOTIME, t2, loc)
  if t1Time.Unix() > t2Time.Unix(){
    return false
  }
  return true
}