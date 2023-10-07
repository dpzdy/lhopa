import Vue from 'vue'
import App from './App.vue'

import axios from 'axios'

Vue.config.productionTip = false

//全局配置axios的请求路径
// axios.defaults.baseURL = 'http://127.0.0.1'
//把axios挂载到vue.prototype上，供每个.vue的实例直接使用
Vue.prototype.$http = axios

new Vue({
  render: h => h(App),
}).$mount('#app')
