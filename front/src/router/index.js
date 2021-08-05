import Vue from 'vue'
import VueRouter from 'vue-router'

const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

const { isNavigationFailure, NavigationFailureType } = VueRouter

const index = () => import(/*webpackMode: "eager" */'@/views/index/index.vue')
const login = () => import(/*webpackMode: "eager" */'@/views/index/login/index.vue')
const dashboard = () => import(/* webpackChunkName: "dashboard", webpackPrefetch: true */ '@/views/dashboard/index.vue')
const playground = () => import(/* webpackChunkName: "dashboard", webpackPrefetch: true */ '@/views/dashboard/content/playground/index.vue')


export const routes = [
    {
        path: '/', component: index, children: [
            {
                path: '', component: login, name: 'login', meta: { keepAlive: true}
            }
        ],
    },
    {
        path: '/dashboard',
        component: dashboard,
        children: [
            {
                path: 'playground',
                component: playground,
                name: 'playground'
            }
        ]
    },
    {
        path: '*',
        redirect: '/'
    }
]

const router = new VueRouter({
    mode: 'history',
    scrollBehavior: () => ({ y: 0 }),  //页面切换时滚动到顶部
    routes
})

router.push('login').catch(failure => {
    console.log("login error");
    if (isNavigationFailure(failure, NavigationFailureType.redirected)) {
        console.log('isNavigationFailure login');
    }
})

const vaultRegex = /^(.*vault.*)|(dashboard\/version_detail)$/

router.beforeEach((to,from,next) => {
   if(vaultRegex.test(from.path) && !vaultRegex.test(to.path)) {
      // navigate from vault to other page
      localStorage.removeItem('VAULT_LOGIN_STATUS')
   }
   next();
})
export default router