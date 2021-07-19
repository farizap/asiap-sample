import authMiddleware from '~auth/core/middleware'
import Middleware from './middleware'
import Auth from '~auth/core/auth'
import ExpiredAuthSessionError from '~auth/inc/expired-auth-session-error'

// Active schemes
import scheme_41b0a53b from '~auth/schemes/oauth2'

Middleware.auth = authMiddleware

export default function (ctx, inject) {
  // Options
  const options = {"resetOnError":false,"scopeKey":"scope","rewriteRedirects":true,"fullPathRedirect":false,"watchLoggedIn":true,"redirect":{"login":"/login","logout":"/","home":"/","callback":"/login"},"vuex":{"namespace":"auth"},"cookie":{"prefix":"auth.","options":{"path":"/"}},"localStorage":{"prefix":"auth."},"defaultStrategy":"keycloak"}

  // Create a new Auth instance
  const $auth = new Auth(ctx, options)

  // Register strategies
  // keycloak
  $auth.registerStrategy('keycloak', new scheme_41b0a53b($auth, {"endpoints":{"authorization":"/auth/realms/asiap/protocol/openid-connect/auth","token":"/auth/realms/asiap/protocol/openid-connect/token","userInfo":"/auth/realms/asiap/protocol/openid-connect/userinfo","logout":"/auth/realms/asiap/protocol/openid-connect/logout?redirect_uri=https%3A%2F%2Flocalhost%3A3000"},"token":{"property":"access_token","type":"Bearer","name":"Authorization","maxAge":300},"refreshToken":{"property":"refresh_token","maxAge":2592000},"responseType":"code","grantType":"authorization_code","clientId":"asiap-client","scope":["openid","profile","email"],"codeChallengeMethod":"S256","name":"keycloak"}))

  // Inject it to nuxt context as $auth
  inject('auth', $auth)
  ctx.$auth = $auth

  // Initialize auth
  return $auth.init().catch(error => {
    if (process.client) {
      // Don't console log expired auth session errors. This error is common, and expected to happen.
      // The error happens whenever the user does an ssr request (reload/initial navigation) with an expired refresh
      // token. We don't want to log this as an error.
      if (error instanceof ExpiredAuthSessionError) {
        return
      }

      console.error('[ERROR] [AUTH]', error)
    }
  })
}
