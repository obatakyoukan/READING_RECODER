import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import BookSearch from '../components/BookSearch.vue'
import BookForm from '../components/BookForm.vue'
import BookMake from '../components/BookMake.vue'
import BookMake2 from '../components/BookMake2.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
   path: '/search',
   name: 'search',
   component: BookSearch
  }, 
  {
   path: '/form',
   name: 'form',
   component: BookForm
  },
  {
   path: '/make',
   name: 'make',
   component: BookMake
  },
  {
   path: '/make2',
   name: 'make2',
   component: BookMake2
  },
  {
   path: '*',
   redirect: '/'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
