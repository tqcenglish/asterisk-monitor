<template>
  <div class="bottom-charts">
    <div class="bc-chart-item">
      <div class="bcci-header">{{ config1.name }}</div>
      <dv-active-ring-chart :config="config1" />
      <Label-Tag :config="labelConfig" />
    </div>
    <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />

    <div class="bc-chart-item">
      <div class="bcci-header">{{ config2.name }}</div>
      <dv-active-ring-chart :config="config2" />
      <Label-Tag :config="labelConfig" />
    </div>
    <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />

    <div class="bc-chart-item">
      <div class="bcci-header">{{ config3.name }}</div>
      <dv-active-ring-chart :config="config3" />
      <Label-Tag :config="labelConfig" />
    </div>
    <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />

    <div class="bc-chart-item">
      <div class="bcci-header">{{ config4.name }}</div>
      <dv-active-ring-chart :config="config4" />
      <Label-Tag :config="labelConfig" />
    </div>
    <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />
    <!-- <div v-for="config in config"  :key="config.name">
      <div class="bc-chart-item">
        <div class="bcci-header">{{config.name}}</div>
        <dv-active-ring-chart :config="config" />
        <Label-Tag :config="labelConfig" />
      </div>
      <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />
    </div> -->

  </div>
</template>

<script>
import LabelTag from './LabelTag'

import { queueStatus } from './api'
export default {
  name: 'Queue',
  components: {
    LabelTag
  },
  data () {
    return {
      config1: {
        radius: '70%',
        activeRadius: '75%',
        data: [
          {
            name: '中止',
            value: 1
          },
          {
            name: '通话中',
            value: 1
          },
          {
            name: '已完成',
            value: 1
          }
        ],
        color: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
        digitalFlopStyle: {
          fontSize: 20
        },
        showOriginValue: true
      },
      config2: {
        radius: '70%',
        activeRadius: '75%',
        data: [
          {
            name: '中止',
            value: 1
          },
          {
            name: '通话中',
            value: 1
          },
          {
            name: '已完成',
            value: 1
          }
        ],
        color: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
        digitalFlopStyle: {
          fontSize: 20
        },
        showOriginValue: true
      },
      config3: {
        radius: '70%',
        activeRadius: '75%',
        data: [
          {
            name: '中止',
            value: 1
          },
          {
            name: '通话中',
            value: 1
          },
          {
            name: '已完成',
            value: 1
          }
        ],
        digitalFlopStyle: {
          fontSize: 20
        },
        showOriginValue: true
      },
      config4: {
        radius: '70%',
        activeRadius: '75%',
        data: [
          {
            name: '中止',
            value: 1
          },
          {
            name: '通话中',
            value: 1
          },
          {
            name: '已完成',
            value: 1
          }
        ],
        digitalFlopStyle: {
          fontSize: 20
        },
        showOriginValue: true
      },
      labelConfig: {
        data: ['中止', '通话中', '已完成']
      }
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
      queueStatus().then((data) => {
        data.forEach((item, index) => {
          if (index > 3) {
            return
          }
          switch (index) {
            case 0:
              const { config1 } = this
              config1.name = `${item.name}(${item.queue})`
              config1.data[0].value = parseInt(item.abandoned)
              config1.data[1].value = parseInt(item.calls)
              config1.data[2].value = parseInt(item.completed)
              if (config1.data[0].value === 0 && config1.data[1].value === 0 && config1.data[2].value === 0) {
                break
              }
              this.config1 = { ...config1 }
              break
            case 1:
              const { config2 } = this
              config2.name = `${item.name}(${item.queue})`
              config2.data[0].value = parseInt(item.abandoned)
              config2.data[1].value = parseInt(item.calls)
              config2.data[2].value = parseInt(item.completed)
              if (config2.data[0].value === 0 && config2.data[1].value === 0 && config2.data[2].value === 0) {
                break
              }
              this.config2 = { ...config2 }
              break
            case 2:
              const { config3 } = this
              config3.name = `${item.name}(${item.queue})`
              config3.data[0].value = parseInt(item.abandoned)
              config3.data[1].value = parseInt(item.calls)
              config3.data[2].value = parseInt(item.completed)
              if (config3.data[0].value === 0 && config3.data[1].value === 0 && config3.data[2].value === 0) {
                break
              }
              this.config3 = { ...config3 }
              break
            case 3:
              const { config4 } = this
              config4.name = `${item.name}(${item.queue})`
              config4.data[0].value = parseInt(item.abandoned)
              config4.data[1].value = parseInt(item.calls)
              config4.data[2].value = parseInt(item.completed)
              if (config4.data[0].value === 0 && config4.data[1].value === 0 && config4.data[2].value === 0) {
                break
              }
              this.config4 = { ...config4 }
              break
            default:
              break
          }
        })
      })
    }
  }
}
</script>

<style lang="less">
.bottom-charts {
  width: 100%;
  height: 100%;
  display: flex;
  position: relative;

  .bc-chart-item {
    width: 25%;
    height: 100%;
    padding-top: 20px;
    box-sizing: border-box;
  }

  .bcci-header {
    height: 50px;
    text-align: center;
    line-height: 50px;
    font-size: 20px;
  }

  .dv-active-ring-chart {
    height: calc(~"100% - 80px");
  }

  .label-tag {
    height: 30px;
  }

  .active-ring-name {
    font-size: 18px !important;
  }

  .decoration-1,
  .decoration-2,
  .decoration-3 {
    display: absolute;
    left: 0%;
  }
}
</style>
