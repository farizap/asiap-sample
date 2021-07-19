import colors from 'vuetify/es5/util/colors'

export default {
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: '%s - client',
    title: 'client',
    htmlAttrs: {
      lang: 'en',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    // https://go.nuxtjs.dev/vuetify
    '@nuxtjs/vuetify',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    ['@nuxtjs/proxy', { ws: false }],
    '@nuxtjs/axios',
    '@nuxtjs/auth-next',
    // '@nuxtjs/auth',
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // baseURL: 'https://localhost:8080',
    // browserBaseURL: 'https://localhost:8080',
    // proxyHeaders: true,
    proxy: true,
  },
  proxy: {
    '/auth': 'https://localhost:8080',
  },
  auth: {
    strategies: {
      local: false,
      keycloak: {
        scheme: 'oauth2',
        endpoints: {
          authorization: '/auth/realms/asiap/protocol/openid-connect/auth',
          token: '/auth/realms/asiap/protocol/openid-connect/token',
          userInfo: '/auth/realms/asiap/protocol/openid-connect/userinfo',
          logout:
            '/auth/realms/asiap/protocol/openid-connect/logout?redirect_uri=' +
            encodeURIComponent('https://localhost:3000'),
        },
        token: {
          property: 'access_token',
          type: 'Bearer',
          name: 'Authorization',
          maxAge: 300,
        },
        refreshToken: {
          property: 'refresh_token',
          maxAge: 60 * 60 * 24 * 30,
        },
        responseType: 'code',
        grantType: 'authorization_code',
        clientId: 'asiap-client',
        scope: ['openid', 'profile', 'email'],
        codeChallengeMethod: 'S256',
      },

      // keycloak: {
      //   scheme: 'oauth2',
      //   endpoints: {
      //     authorization: '/auth/realms/asiap/protocol/openid-connect/auth',
      //     token: '/auth/realms/asiap/protocol/openid-connect/token',
      //     userInfo: '/auth/realms/asiap/protocol/openid-connect/userinfo',
      //     logout:
      //       '/auth/realms/asiap/protocol/openid-connect/logout?redirect_uri=' +
      //       encodeURIComponent('https://localhost:3000'),
      //   },
      //   token: {
      //     property: 'access_token',
      //     type: 'Bearer',
      //     name: 'Authorization',
      //     maxAge: 300,
      //   },
      //   refreshToken: {
      //     property: 'refresh_token',
      //     maxAge: 60 * 60 * 24 * 30,
      //   },
      //   responseType: 'code',
      //   grantType: 'authorization_code',
      //   clientId: 'asiap-client',
      //   scope: ['openid', 'profile', 'email'],
      //   codeChallengeMethod: 'S256',
      // },
    },
    redirect: {
      login: '/login',
      logout: '/',
      home: '/',
    },
  },
  router: {
    middleware: ['auth'],
  },

  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: true,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3,
        },
      },
    },
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
}
