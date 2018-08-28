import Vue from 'vue'
import Router from 'vue-router'
import index from '@/views/index'
import json from '@/views/tools/json'
import category from '@/views/category'

Vue.use(Router)

export default new Router({
  routes: [
    { path: '/', name: 'index', component: index },
    { path: '/c/:code', name: 'category', component: category },
    { path: '/json', name: 'json', component: json }
  ]
})
