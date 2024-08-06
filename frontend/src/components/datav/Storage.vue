<template>
  <div class="left-chart-2">
    <div class="lc2-header">存储信息</div>
    <div class="lc2-details">设备空间<span>245 GB</span></div>
    <dv-charts class="lc2-chart" :option="option" />
    <dv-decoration-2 style="height:10px;" />
  </div>
</template>

<script>
import { storageInfo } from './api'
export default {
  name: 'LeftChart2',
  data () {
    return {
      option: {
        series: [
          {
            type: 'pie',
            data: [
              { name: '通话录音', value: 93 },
              { name: '系统', value: 32 },
              { name: '语音邮件', value: 65 },
              { name: '音乐文件', value: 44 },
              { name: '其他', value: 52 }
            ],
            radius: ['45%', '65%'],
            insideLabel: {
              show: false
            },
            outsideLabel: {
              labelLineEndLength: 10,
              formatter: '{percent}%\n{name}',
              style: {
                fontSize: 14,
                fill: '#fff'
              }
            }
          }
        ],
        color: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b']
      }
    }
  },
  created: function () {
    storageInfo().then((data) => {
      let { option } = this
      option.series[0].data[0].value = data.record
      option.series[0].data[1].value = data.system
      option.series[0].data[2].value = data.voicemail
      option.series[0].data[3].value = data.music
      option.series[0].data[4].value = data.other
      this.option = { ...option }
    })
  }
}
</script>

<style lang="less">
.left-chart-2 {
  margin-top: 30px;
  width: 100%;
  height: 30%;
  display: flex;
  flex-direction: column;

  .lc2-header {
    height: 20px;
    line-height: 20px;
    font-size: 16px;
    text-indent: 20px;
    margin-top: 10px;
  }

  .lc2-details {
    height: 40px;
    font-size: 16px;
    display: flex;
    align-items: center;
    text-indent: 20px;

    span {
      color: #096dd9;
      font-weight: bold;
      font-size: 35px;
      margin-left: 20px;
    }
  }

  .lc2-chart {
    height: calc(~"100% - 80px");
  }
}
</style>
