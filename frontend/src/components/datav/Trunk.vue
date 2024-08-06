<template>
  <div class="right-chart-1">
    <div class="rc1-header">中继累计通话</div>

    <div class="rc1-chart-container">
      <div class="left">
        <div class="number">{{ total }}</div>
        <div>线路总数</div>
      </div>

      <dv-capsule-chart class="right" :config="config" />
    </div>
  </div>
</template>

<script>
import { trunkStatus } from './api'
export default {
  name: 'Trunk',
  data () {
    return {
      config: {
        data: [
          {
            name: '告警线路',
            value: 25
          },
          {
            name: '客服线路',
            value: 66
          }
        ]
        // unit: '路'
      },
      total: 0
    }
  },
  created: function () {
    this.getData()
    this.timer = setInterval(() => {
      this.getData()
    }, 10000)
  },
  beforeDestroy () {
    clearInterval(this.timer)
  },
  methods: {
    getData: function () {
      trunkStatus().then((data) => {
        let keys = Object.keys(data)
        this.total = keys.length
        this.config.data = keys.map(item => {
          return {
            name: item,
            value: data[item]
          }
        })
        this.config = { ...this.config }
      })
    }
  }
}
</script>

<style lang="less">
.right-chart-1 {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;

  .rc1-header {
    padding: 10px;
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
