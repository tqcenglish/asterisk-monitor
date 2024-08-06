<template>
  <div class="left-chart-3">
    <div class="lc3-header">网络信息</div>
    <div class="lc3-details">阻止攻击<span>215 次</span></div>
    <dv-capsule-chart class="lc3-chart" :config="config" />
  </div>
</template>

<script>
import { networkInfo } from './api'
export default {
  name: 'Network',
  data () {
    return {
      config: {
        data: [
          {
            name: '发送流量',
            value: 1
          },
          {
            name: '接收流量',
            value: 1
          }
        ],
        colors: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
        unit: 'Gb'
      }
    }
  },

  created: function () {
    this.getData()
    this.timer = setInterval(() => {
      this.getData()
    }, 60000)
  },
  beforeDestroy () {
    clearInterval(this.timer)
  },
  methods: {
    getData: function () {
      networkInfo().then((data) => {
        let { config } = this
        config.data[0].value = data['tx']
        config.data[1].value = data['rx']
        console.log(data)
        this.config = { ...config }
      })
    }
  }
}
</script>

<style lang="less">
.left-chart-3 {
  margin-top: 30px;
  width: 100%;
  height: 33%;
  display: flex;
  flex-direction: column;

  .lc3-header {
    height: 20px;
    line-height: 20px;
    font-size: 16px;
    text-indent: 20px;
    margin-top: 10px;
  }

  .lc3-details {
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

  .lc3-chart {
    flex: 1;
  }
}
</style>
