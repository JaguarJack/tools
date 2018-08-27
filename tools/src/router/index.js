import Vue from 'vue'
import Router from 'vue-router'
import index from '@/components/index'
import json from '@/components/tools/json'
Vue.use(Router)

export default new Router({
  routes: [
    { path: '/c/:code', name: 'index', component: index },
    { path: '/json', name: 'json', component: json }
  ]
})
