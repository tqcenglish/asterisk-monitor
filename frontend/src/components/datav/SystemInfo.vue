<template>
  <div class="left-chart-1">
    <div class="lc1-header">设备信息</div>
    <div class="lc1-details">无故障运行时间<span>99 小时 </span></div>
    <dv-capsule-chart class="lc1-chart" :config="config" />
    <dv-decoration-2 style="height:10px;" />
  </div>
</template>

<script>
export default {
  name: 'SystemInfo',
  data () {
    return {
      config: {
        data: [
          {
            name: 'Web 服务时间',
            value: 10
          },
          {
            name: 'Asterisk 服务时间',
            value: 10
          },
          {
            name: '系统启动时间',
            value: 10
          }
        ],
        // colors: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
        colors: ['#00baff', '#3de7c9', '#ffc530'],
        unit: '小时'
      }
    }
  },
  created: function () {
    fetch('http://127.0.0.1:8080/api/systeminfo').then((res) => {
      res.json().then((data) => {
        this.config.data[0].value = data.webUptime
        this.config.data[1].value = data.voipUptime
        this.config.data[2].value = data.systemUptime
        this.config = { ...this.config }
      })
    })
  }
}
</script>

<style lang="less">
.left-chart-1 {
  width: 100%;
  height: 37%;
  display: flex;
  flex-grow: 0;
  flex-direction: column;

  .lc1-header {
    text-align: center;
    height: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 30px;
    margin-bottom: 20px;
  }

  .lc1-details {
    height: 50px;
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

  .lc1-chart {
    flex: 1;
  }
}
</style>
