// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import iView from 'iview'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Qs from 'qs'
import 'iview/dist/styles/iview.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'

var _axios = axios.create({
  transformRequest: [function (data) {
    return Qs.stringify(data)
  }],
  headers: {'Content-Type': 'application/x-www-form-urlencoded'}
})

Vue.use(VueAxios, _axios)
Vue.use(iView)

Vue.config.productionTip = false
Vue.prototype.http = _axios
Vue.prototype.host = 'http://192.168.36.9:9092/'

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
