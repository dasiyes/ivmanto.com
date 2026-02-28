import { computed, watch } from 'vue'
import { useRoute } from 'vue-router'

//   '/services/go-backend-development': {
//   title: 'Go Backend Development | ivmanto.com',
//   description:
//     'High-performance Go (Golang) backend development for data-intensive applications. We build scalable, concurrent, and efficient cloud-native services.',
// },

// SEO metadata mapping for specific routes
const routeMetadata: Record<string, { title: string; description: string }> = {
  '/': {
    title: 'ivmanto.com | Data & AI Consultancy',
    description:
      'Expert Data & AI consultancy specializing in Google Cloud Platform (GCP). We help businesses with data architecture, governance, and AI-driven solutions to turn data into a strategic asset.',
  },
  '/services': {
    title: 'Services | ivmanto.com',
    description:
      'Explore our Data & AI services. From data strategy and GCP architecture to custom AI/ML solutions and Go backend development, we empower your business with data.',
  },
  '/services/data-strategy-and-governance': {
    title: 'Data Strategy & Governance | ivmanto.com',
    description:
      'Develop a clear data strategy and robust governance framework. We align your data initiatives with business goals for maximum impact and compliance.',
  },
  '/services/data-architecture': {
    title: 'Data Architecture on GCP | ivmanto.com',
    description:
      'Design and build scalable, secure data architectures on Google Cloud Platform (GCP). We leverage BigQuery, Cloud Storage, and modern data engineering practices.',
  },
  '/services/sovereigncloud': {
    title: 'Sovereign Cloud Solutions | ivmanto.com',
    description:
      'Explore architectural perspectives on Data, Operations, and AI Sovereignty to meet your compliance and security needs in the cloud.',
  },
  '/services/ml-engineering': {
    title: 'AI & ML Solutions | ivmanto.com',
    description:
      'Leverage the power of AI and Machine Learning on GCP. We build custom solutions, from predictive analytics to generative AI, to solve your toughest challenges.',
  },
  '/services/principles': {
    title: 'Guiding Principles | ivmanto.com',
    description:
      'Our DAMA-aligned principles for data strategy, governance, and architecture ensure your data becomes a reliable, valuable asset for decision-making and AI.',
  },
  '/blog': {
    title: 'Insights & Articles | ivmanto.com',
    description:
      'Read our latest articles and insights on data strategy, cloud architecture, AI/ML, and software engineering. Stay ahead of the curve with expert analysis.',
  },
  '/about': {
    title: 'About | ivmanto.com',
    description:
      'Learn about IVMANTO and our mission to help businesses harness the power of data. Meet the experts behind our innovative data and AI solutions.',
  },
  '/booking': {
    title: 'Contact us | ivmanto.com',
    description:
      'Get in touch with IVMANTO to discuss your data and AI challenges. Book a free consultation or send us a message to start your data transformation journey.',
  },
  '/privacy-policy': {
    title: 'Privacy Policy | ivmanto.com',
    description:
      'Read the IVMANTO Privacy Policy to understand how we collect, use, and protect your personal data in accordance with GDPR and other regulations.',
  },
}

// Default metadata for other pages
const defaultTitle = 'ivmanto.com | Data & AI Consultancy'
const defaultDescription =
  'Expert Data & AI consultancy specializing in Google Cloud Platform (GCP). We help businesses with data architecture, governance, and AI-driven solutions to turn data into a strategic asset.'

export function usePageMetadata() {
  const route = useRoute()

  // Dynamically computed metadata based on the current route
  const pageTitle = computed(() => routeMetadata[route.path]?.title ?? defaultTitle)
  const pageDescription = computed(
    () => routeMetadata[route.path]?.description ?? defaultDescription,
  )
  // 👇 ADD THIS NEW COMPUTED PROPERTY 👇
  const cleanTitle = computed(() => pageTitle.value.split(' | ')[0])

  return { pageTitle, pageDescription, cleanTitle }
}
