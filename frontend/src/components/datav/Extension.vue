<template>
  <div class="right-chart-2">
    <div class="rc1-header">分机状态</div>

    <div class="rc1-chart-container">
      <div class="left">
        <div class="number">{{ total }}</div>
        <div>分机总数</div>
      </div>

      <dv-charts class="right" :option="option" />
    </div>
  </div>
</template>

<script>
import { extensionStatus } from './api'
export default {
  name: 'Extension',
  data () {
    return {
      option: {
        series: [
          {
            type: 'pie',
            data: [
              { name: '在线', value: 1 },
              { name: '通话中', value: 1 },
              { name: '离线', value: 1 }
            ],
            radius: ['45%', '65%'],
            insideLabel: {
              show: false
            },
            outsideLabel: {
              labelLineEndLength: 20,
              formatter: '{value}\n{name}',
              style: {
                fontSize: 14,
                fill: '#fff'
              }
            }
          }
        ],
        color: ['#00baff', '#3de7c9', '#ffc530', '#469f4b']
      },
      total: 0
    }
  },
  created: function () {
    this.getData()
    this.timer = setInterval(() => {
      this.getData()
    }, 3000)
  },
  beforeDestroy () {
    clearInterval(this.timer)
  },
  methods: {
    getData: function () {
      extensionStatus().then((data) => {
      // Busy
      // In use
      // Not in use
      // On Hold
      // Ringing
      // Unavailable
        // console.log(data)
        this.option.series[0].data[0].value = data['Not in use']
        this.option.series[0].data[1].value = data['Busy'] + data['In use'] + data['Ringing']
        this.option.series[0].data[2].value = data['Unavailable']

        this.option = { ...this.option }
        this.total = data['Not in use'] + data['Busy'] + data['In use'] + data['Ringing'] + data['Unavailable']
        console.log(this.option)
      })
    }
  }
}
</script>

<style lang="less">
.right-chart-2 {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;

  .rc1-header {
    font-size: 24px;
    font-weight: bold;
    height: 30px;
    line-height: 30px;
  }

  .rc1-chart-container {
    flex: 1;
    display: flex;
  }

  .left {
    width: 30%;
    font-size: 16px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    .number {
      font-size: 34px;
      color: #096dd9;
      font-weight: bold;
      margin-bottom: 30px;
    }
  }

  .right {
    flex: 1;
    padding-bottom: 20px;
    padding-right: 20px;
    box-sizing: border-box;
  }
}
</style>
