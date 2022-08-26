import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import installElementPlus from './plugins/element'
import './assets/css/global.css'
import axios from 'axios'
import "./assets/global.css"
import echarts from 'echarts'

axios.defaults.baseURL = "http://localhost:8080/"
const app = createApp(App)

app.config.globalProperties.$echarts = echarts
app.config.globalProperties.$http = axios
installElementPlus(app)
app.use(router).mount('#app')
