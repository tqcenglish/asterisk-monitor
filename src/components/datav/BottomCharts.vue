<template>
  <div class="bottom-charts">
    <div class="bc-chart-item">
      <div class="bcci-header">销售队列</div>
      <dv-active-ring-chart :config="config1" />
      <Label-Tag :config="labelConfig" />
    </div>
    <dv-decoration-2 class="decoration-1" :reverse="true" style="width:5px;" />

    <div class="bc-chart-item">
      <div class="bcci-header">技术支持</div>
      <dv-active-ring-chart :config="config2" />
      <Label-Tag :config="labelConfig" />
    </div>

    <dv-decoration-2 class="decoration-2" :reverse="true" style="width:5px;" />

    <div class="bc-chart-item">
      <div class="bcci-header">研发支持</div>
      <dv-active-ring-chart :config="config3" />
      <Label-Tag :config="labelConfig" />
    </div>

    <dv-decoration-2 class="decoration-3" :reverse="true" style="width:5px;" />

    <div class="bc-chart-item">
      <div class="bcci-header">客户队列</div>
      <dv-active-ring-chart :config="config4" />
      <Label-Tag :config="labelConfig" />
    </div>

  </div>
</template>

<script>
import LabelTag from './LabelTag'

export default {
  name: 'BottomCharts',
  components: {
    LabelTag
  },
  data () {
    return {
      config1: {
        data: [
          {
            name: '销售',
            value: 356
          },
          {
            name: '技术',
            value: 215
          },
          {
            name: '研发',
            value: 90
          },
          {
            name: '客户',
            value: 317
          }
        ],
        color: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
        radius: '65%',
        activeRadius: '70%'
      },

      config2: {
        data: [
          {
            name: '收费站',
            value: 615
          },
          {
            name: '监控中心',
            value: 322
          },
          {
            name: '道路外场',
            value: 198
          },
          {
            name: '其他',
            value: 80
          }
        ],
        color: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
        radius: '65%',
        activeRadius: '70%'
      },

      config3: {
        data: [
          {
            name: '收费站',
            value: 452
          },
          {
            name: '监控中心',
            value: 512
          },
          {
            name: '道路外场',
            value: 333
          },
          {
            name: '其他',
            value: 255
          }
        ],
        color: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
        radius: '65%',
        activeRadius: '70%'
      },

      config4: {
        data: [
          {
            name: '收费站',
            value: 156
          },
          {
            name: '监控中心',
            value: 415
          },
          {
            name: '道路外场',
            value: 90
          },
          {
            name: '其他',
            value: 210
          }
        ],
        color: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
        radius: '65%',
        activeRadius: '70%'
      },

      labelConfig: {
        data: ['呼入', '呼出', '内部', '其他']
      }
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
