<template>
	<div>
		<card></card>
		<h1>堆存费可视化</h1>
		<div class="center-block">
			<div class="echart" id="mychart" :style="myChartStyle">
			</div>
		</div>
	</div>

</template>

<script>
import * as echarts from 'echarts'
// import axios from 'axios'
export default {
	data() {
		return {
			xData: [], //横坐标
			yData: [], //数据
			myChartStyle: {
				float: 'left',
				width: '80%',
				height: '400px'
			} //图表样式
		}
	},
	mounted() {

		this.getview()


	},

	methods: {
		async getview() {
			const { data: res } = await this.$http.get('guobaichuan/route/view')

			this.xData = res.data.x
			this.yData = res.data.y
			this.initEcharts()
			console.log(this.xData)

		},
		async initEcharts() {



			// 基本柱状图
			const option = {
				tooltip: {
					trigger: 'axis',
					axisPointer: {
						type: 'shadow'
					}
				},
				grid: {
					containLabel: true
				},
				color: ['#409EFF'],
				title: {
					text: '堆存费表',
					left: 'center'
				},
				xAxis: {
					name: '单号',
					data: this.xData,
					axisTick: {
						alignWithLabel: true
					}
				},
				yAxis: {},
				series: [{
					name: '费用',
					type: 'bar', //形状为柱状图
					data: this.yData
					// markLine: {
					// 	data: [{
					// 		type: 'average',
					// 		name: '平均值'
					// 	}]
					// }

				}]
			}
			const myChart = echarts.init(document.getElementById('mychart'))
			myChart.setOption(option, true)
			//随着屏幕大小调节图表
			window.addEventListener('resize', () => {
				myChart.resize()
			})
		}
	}
}
</script>
<style>
.center-block {
	margin: 0 auto;
	width: 90%;
	padding: 1rem;
	color: #fff;

}
</style>
