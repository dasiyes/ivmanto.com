import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import AboutView from '@/views/AboutView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/services',
      name: 'services',
      component: () => import('../views/ServiceView.vue'),
      children: [
        {
          path: '',
          name: 'services-index',
          component: () => import('../components/sections/ServicesIndex.vue'),
        },
        {
          path: ':id',
          name: 'service-detail',
          component: () => import('../components/sections/ServicesSection.vue'),
          props: true, // Pass route params as props to the component
        },
      ],
    },
    {
      path: '/blog',
      name: 'blog',
      // This is a wrapper component for the blog section
      component: () => import('../views/BlogView.vue'),
      children: [
        {
          path: '',
          name: 'blog-index',
          component: () => import('../views/BlogIndexView.vue'),
        },
        {
          path: ':slug',
          name: 'blog-post',
          component: () => import('../views/BlogPostView.vue'),
          props: true, // Pass route params as props
        },
      ],
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
  ],
  scrollBehavior(to) {
    if (to.hash) {
      return { el: to.hash, behavior: 'smooth' }
    }
    return { top: 0 }
  },
})

export default router
