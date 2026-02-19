import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import AboutView from '@/views/AboutView.vue'
import BlogView from '@/views/BlogView.vue'
import ArticleListView from '@/views/ArticleListView.vue'
import ArticleView from '@/views/ArticleView.vue'
import ServiceView from '@/views/ServiceView.vue' // This will be our new dynamic view
import ServicesLanding from '@/views/ServicesLanding.vue'
import BookingCalendar from '@/views/BookingCalendar.vue'
import { services } from '@/data/services'
import BookingCancellation from '@/views/BookingCancellation.vue'
import BookingGoogleDemo from '@/views/BookingGoogleDemo.vue'
import NotFoundView from '@/views/NotFoundView.vue' // Import 404 View

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
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
          beforeEnter: (to, from, next) => {
            // Check if slug starts with an underscore (e.g. _DamaPrinciples)
            if (to.params.slug && typeof to.params.slug === 'string' && to.params.slug.startsWith('_')) {
              // Redirect to the clean slug (remove leading underscore)
              next({ name: 'article-detail', params: { slug: to.params.slug.substring(1) }, replace: true });
            } else {
              next();
            }
          }
        },
      ],
    },
    {
      // This route handles the generic '/services' path.
      // It restores the `name: 'services'` route used by the "More About our Services" button on the Home page.
      // It then redirects to the first available service page.
      path: '/services',
      name: 'services',
      component: ServicesLanding,
    },
    {
      // This single dynamic route handles all service pages.
      // e.g., /services/data-architecture
      path: '/services/:id',
      name: 'service-detail',
      component: () => import('@/views/ServiceView.vue'),
      props: true, // Passes the :id from the URL as a prop to ServiceView
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView,
    },
    {
      path: '/booking', // <-- Add this new route object
      name: 'booking',
      component: BookingCalendar,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/booking/cancel',
      name: 'BookingCancellation',
      component: BookingCancellation,
    },
    {
      path: '/booking-demo',
      name: 'booking-demo',
      component: BookingGoogleDemo,
    },
    {
      path: '/privacy-policy',
      name: 'privacy-policy',
      // This uses lazy-loading for better performance
      component: () => import('../components/PrivacyPolicy.vue'),
    },
    // Catch-All 404 Route
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: NotFoundView,
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
