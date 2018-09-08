import Vue from 'vue'
import Router from 'vue-router'
import index from '@/views/index'
import json from '@/views/tools/json'
import ssh from '@/views/tools/ssh'
import category from '@/views/category'
import gif from '@/views/tools/gif'

Vue.use(Router)

export default new Router({
  routes: [
    { path: '/', name: 'index', component: index },
    { path: '/c/:code', name: 'category', component: category },
    { path: '/json', name: 'json', component: json },
    { path: '/ssh', name: 'ssh', component: ssh },
    { path: '/gif', name: 'gif', component: gif }
  ]
})
