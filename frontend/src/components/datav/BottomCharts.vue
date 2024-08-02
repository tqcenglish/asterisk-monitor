<template>
  <div class="bottom-charts">
    <div class="bc-chart-item" v-if="configs[0]">
      <div class="bcci-header">{{ configs[0].name }}</div>
      <dv-active-ring-chart :config="configs[0]" />
      <Label-Tag :config="labelConfig" />
    </div>
    <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />

    <div class="bc-chart-item" v-if="configs[1]">
      <div class="bcci-header">{{ configs[1].name }}</div>
      <dv-active-ring-chart :config="configs[1]" />
      <Label-Tag :config="labelConfig" />
    </div>
    <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />

    <div class="bc-chart-item" v-if="configs[2]">
      <div class="bcci-header">{{ configs[2].name }}</div>
      <dv-active-ring-chart :config="configs[2]" />
      <Label-Tag :config="labelConfig" />
    </div>
    <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />

    <div class="bc-chart-item" v-if="configs[3]">
      <div class="bcci-header">{{ configs[3].name }}</div>
      <dv-active-ring-chart :config="configs[3]" />
      <Label-Tag :config="labelConfig" />
    </div>
    <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />

    <!-- <div v-for="config in configs"  :key="config.name">
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
  name: 'BottomCharts',
  components: {
    LabelTag
  },
  data () {
    return {
      configs: [],
      labelConfig: {
        data: ['中止', '通话中', '已完成']
      }
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
      queueStatus().then((data) => {
        this.configs = data.params
          .filter((_, index) => {
            if (index > 4) {
              return false
            }
            return true
          })
          .map(item => {
            return {
              name: item.queue,
              data: [
                {
                  name: '中止',
                  value: parseInt(item.abandoned) === 0 ? 1 : parseInt(item.abandoned)
                },
                {
                  name: '通话中',
                  value: parseInt(item.calls) === 0 ? 1 : parseInt(item.calls)
                },
                {
                  name: '已完成',
                  value: parseInt(item.completed) === 0 ? 1 : parseInt(item.completed)
                }
              ],
              color: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
              lineWidth: 30,
              radius: '65%',
              activeRadius: '70%',
              digitalFlopStyle: {
                fontSize: 20
              },
              showOriginValue: true
            }
          })
        console.log(this.configs)
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

  .decoration-1, .decoration-2, .decoration-3 {
    display: absolute;
    left: 0%;
  }
}
</style>
