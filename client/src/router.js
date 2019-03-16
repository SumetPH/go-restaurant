import Vue from 'vue'
import Router from 'vue-router'
import Index from './views/Index.vue'
import Store from './views/Store.vue'
import Search from './views/Search.vue'
import Edit from './views/Edit.vue'
import Delete from './views/Delete.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/store',
      name: 'store',
      component: Store
    },
    {
      path: '/search',
      name: 'search',
      component: Search
    },
    {
      path: '/edit',
      name: 'edit',
      component: Edit
    },
    {
      path: '/delete',
      name: 'delete',
      component: Delete
    }
  ]
})
