import Vue from 'vue'
import Router from 'vue-router'
import index from '@/components/index'
import json from '@/components/tools/json'
import category from '@/components/category'
import slider from '@/components/public/slider'

Vue.use(Router)

export default new Router({
  routes: [
    { path: '/', name: 'index', component: index },
    { path: '/c/:code', name: 'category', component: category },
    { path: '/json', name: 'json', component: {default: json, slider: slider} }
  ]
})
