import Vue from 'vue'
import Router from 'vue-router'
import { normalizeURL, decode } from 'ufo'
import { interopDefault } from './utils'
import scrollBehavior from './router.scrollBehavior.js'

const _6b4dc77e = () => interopDefault(import('..\\pages\\inspire.vue' /* webpackChunkName: "pages/inspire" */))
const _cbe73de0 = () => interopDefault(import('..\\pages\\login.vue' /* webpackChunkName: "pages/login" */))
const _5f36840e = () => interopDefault(import('..\\pages\\index.vue' /* webpackChunkName: "pages/index" */))
const _4d119c84 = () => interopDefault(import('..\\pages\\_continent\\_mountain.vue' /* webpackChunkName: "pages/_continent/_mountain" */))

const emptyFn = () => {}

Vue.use(Router)

export const routerOptions = {
  mode: 'history',
  base: '/',
  linkActiveClass: 'nuxt-link-active',
  linkExactActiveClass: 'nuxt-link-exact-active',
  scrollBehavior,

  routes: [{
    path: "/inspire",
    component: _6b4dc77e,
    name: "inspire"
  }, {
    path: "/login",
    component: _cbe73de0,
    name: "login"
  }, {
    path: "/",
    component: _5f36840e,
    name: "index"
  }, {
    path: "/:continent/:mountain?",
    component: _4d119c84,
    name: "continent-mountain"
  }],

  fallback: false
}

export function createRouter (ssrContext, config) {
  const base = (config._app && config._app.basePath) || routerOptions.base
  const router = new Router({ ...routerOptions, base  })

  // TODO: remove in Nuxt 3
  const originalPush = router.push
  router.push = function push (location, onComplete = emptyFn, onAbort) {
    return originalPush.call(this, location, onComplete, onAbort)
  }

  const resolve = router.resolve.bind(router)
  router.resolve = (to, current, append) => {
    if (typeof to === 'string') {
      to = normalizeURL(to)
    }
    return resolve(to, current, append)
  }

  return router
}
