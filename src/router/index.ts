import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import AboutView from '@/views/AboutView.vue'
import BlogView from '@/views/BlogView.vue'
import ArticleListView from '@/views/ArticleListView.vue'
import ArticleView from '@/views/ArticleView.vue'

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
      // This is the parent route for the blog section
      path: '/blog',
      component: BlogView,
      children: [
        {
          path: '', // Renders ArticleListView at /blog
          name: 'blog',
          component: ArticleListView,
        },
        {
          path: ':slug', // Renders ArticleView at /blog/your-slug
          name: 'article-detail',
          component: ArticleView,
          props: true, // This passes the ':slug' as a prop to ArticleView
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
