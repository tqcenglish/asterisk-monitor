<template>
  <div class="center-cmp">
    <div class="cc-header">
      <dv-decoration-1 style="width:200px;height:50px;" />
      <div>今日通话统计</div>
      <dv-decoration-1 style="width:200px;height:50px;" />
    </div>

    <div class="cc-details">
      <div>通话总数</div>
      <div class="card" v-bind:key="data" v-for="data in total.toString()">
        {{ data }}
      </div>
    </div>

    <div class="cc-main-container">
      <div class="ccmc-left">
        <div class="station-info">
          呼入<span style="width: 40px;">{{ incomingCallCount }}</span>
        </div>
        <div class="station-info">
          呼出<span style="width: 40px;">{{ outgoingCallCount }}</span>
        </div>
        <div class="station-info">
          内部<span style="width: 40px;">{{ internalCallCount }}</span>
        </div>
      </div>

      <dv-active-ring-chart class="ccmc-middle" :config="config" />

      <!-- <div class="ccmc-right">
        <div class="station-info">
          <span>{{ internalCallCount }}</span>内部
        </div>
      </div> -->

      <LabelTag :config="labelConfig" />
    </div>
  </div>
</template>

<script>
import LabelTag from './LabelTag'
import { callLog } from './api'

export default {
  name: 'CenterCmp',
  components: {
    LabelTag
  },
  data () {
    return {
      config: {
        data: [
          {
            name: '呼入',
            value: 10
          },
          {
            name: '呼出',
            value: 10
          },
          {
            name: '内部',
            value: 10
          }
        ],
        color: ['#00baff', '#3de7c9', '#fff', '#ffc530', '#469f4b'],
        lineWidth: 30,
        radius: '55%',
        activeRadius: '60%'
      },
      labelConfig: {
        data: ['呼入', '呼出', '内部']
      },
      total: 0,
      incomingCallCount: 0,
      outgoingCallCount: 0,
      internalCallCount: 0

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
    getData () {
      callLog().then((data) => {
        const { config } = this
        config.data[0].value = data.incomingCallCount
        config.data[1].value = data.outgoingCallCount
        config.data[2].value = data.internalCallCount
        this.config = { ...config }

        this.incomingCallCount = data.incomingCallCount
        this.outgoingCallCount = data.outgoingCallCount
        this.internalCallCount = data.internalCallCount
        this.total = data.internalCallCount + data.incomingCallCount + data.outgoingCallCount
      })
    }
  }
}

</script>

<style lang="less">
.center-cmp {
  width: 100%;
  height: 100%;
  margin: 0px;
  padding: 0px;
  display: flex;
  flex-direction: column;

  .cc-header {
    height: 70px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 30px;
  }

  .cc-details {
    height: 120px;
    display: flex;
    justify-content: center;
    font-size: 32px;
    align-items: center;

    .card {
      background-color: rgba(4,49,128,.6);
      color: #08e5ff;
      height: 70px;
      width: 70px;
      font-size: 45px;
      font-weight: bold;
      line-height: 70px;
      text-align: center;
      margin: 10px;
    }
  }

  .cc-main-container {
    position: relative;
    flex: 1;
    display: flex;

    .ccmc-middle {
      width: 50%;
      height: 90%;

      .active-ring-name {
        font-size: 20px !important;
      }
    }

    .ccmc-left, .ccmc-right {
      width: 25%;
      display: flex;
      flex-direction: column;
      justify-content: center;
      font-size: 24px;

      span {
        font-size: 40px;
        font-weight: bold;
      }

      .station-info {
        height: 80px;
        display: flex;
        align-items: center;
      }
    }

    .ccmc-left {
      align-items: flex-end;

      span {
        margin-left: 20px;
      }
    }

    .ccmc-right {
      align-items: flex-start;

      span {
        margin-right: 20px;
      }
    }
  }

  .label-tag {
    position: absolute;
    width: 500px;
    height: 30px;
    bottom: 10px;
    left: 50%;
    transform: translateX(-50%);
  }
}
</style>
