import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/page/Home.vue'
import Login from '../components/page/Login.vue'
import Posts from "../components/page/Posts.vue"
import NewPost from "../components/page/NewPost.vue"
import UpdatePost from "../components/page/UpdatePost.vue"
import Post from "../components/page/Post.vue"

Vue.use(VueRouter)

const requireAuth = () => (from, to, next) => {
  console.log(from.headers.get('Authorization'));
  const isAuthenticated = false

  if (isAuthenticated) return next()
  next({
    path: '/login',
    query: { redirect: from.fullPath }
  })
}

  const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/posts',
    name: 'Posts',
    component: Posts
  },
  {
    path: '/new/post',
    name: 'NewPost',
    component: NewPost,
    beforeEnter: requireAuth()
  },
  {
    path: '/post/:id/update',
    name: 'UpdatePost',
    component: UpdatePost
  },
  {
    path: '/post/:id',
    name: 'Post',
    component: Post
  }
]

const router = new VueRouter({
  mode:'history',
  routes
})

export default router
