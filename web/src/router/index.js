import Vue from 'vue'
import VueRouter from 'vue-router'

const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push(location) {
   return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

const {
   isNavigationFailure,
   NavigationFailureType
} = VueRouter

const index = () => import( /*webpackMode: "eager" */ '@/index.vue')
const login = () => import( /*webpackMode: "eager" */ '@/views/login/login.vue')
const dashboard = () => import( /* webpackChunkName: "dashboard", webpackPrefetch: true */ '@/views/dashboard/index.vue')
const playground = () => import( /* webpackChunkName: "playground", webpackPrefetch: true */ '@/views/dashboard/content/playground/index.vue')
const userlist = () => import( /* webpackChunkName: "userlist", webpackPrefetch: true */ '@/views/dashboard/content/userlist/index.vue')
const myTasks = () => import( /* webpackChunkName: "myTasks", webpackPrefetch: true */ '@/views/dashboard/content/myTasks/index.vue')
const node = () => import( /* webpackChunkName: "node", webpackPrefetch: true */ '@/views/dashboard/content/node/index.vue')
const profile = () => import( /* webpackChunkName: "changepwd", webpackPrefetch: true */ '@/views/dashboard/content/profile/index.vue')
const taskDetail = () => import( /* webpackChunkName: "taskDetail", webpackPrefetch: true */ '@/views/dashboard/content/taskdetail/index.vue')
const postRegist = () => import( /* webpackChunkName: "navbar", webpackPrefetch: true */ '@/views/dashboard/content/postRegist/index.vue')
export const routes = [{
      path: '/',
      component: index,
      name: 'index',
      children: [{
            path: '/',
            name: 'login',
            component: login
         },
         {
            path: 'dashboard',
            name: 'dashboard',
            component: dashboard,
            children: [{
                  path: 'playground',
                  name: 'playground',
                  component: playground,
               },
               {
                  path: 'myTasks',
                  component: myTasks,
                  name: 'myTasks'
               },

               {
                  path: 'userlist',
                  component: userlist,
                  name: 'userlist'
               },
               {
                  path: 'task/:task_id',
                  component: taskDetail,
                  name: 'taskDetail'
               },
               {
                  path: 'listnodes',
                  component: node,
                  name: 'listnodes'
               },
               {
                  path: 'profile',
                  component: profile,
                  name: 'profile'
               }
            ]
         },
         {
            path: 'post-regist',
            component: postRegist,
            name: 'post-regist'
         },
      ]
   }, 
   {
      path: '*',
      redirect: '/'
   }
]

const router = new VueRouter({
   mode: 'history',
   scrollBehavior: () => ({
      y: 0
   }), //页面切换时滚动到顶部
   routes
})

router.push('login').catch(failure => {
   console.log("login error");
   if (isNavigationFailure(failure, NavigationFailureType.redirected)) {
      console.log('isNavigationFailure login');
   }
})
export default router