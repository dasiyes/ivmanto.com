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
    title: 'Freelance Cloud Data & AI Agent Engineer for SMBs | ivmanto.com',
    description:
      'Hands-on cloud data engineering, custom AI agents and private RAG knowledge systems for small and medium businesses. Senior independent engineer based in Germany, serving the EU and clients worldwide.',
  },
  // Expertise pages render long-form components that read cleanTitle from here
  '/services/data-strategy-and-governance': {
    title: 'Data Strategy & Governance | ivmanto.com',
    description:
      'A clear data strategy and practical governance framework that aligns your data initiatives with business goals, compliance and AI readiness.',
  },
  '/services/data-architecture': {
    title: 'Data Architecture on GCP | ivmanto.com',
    description:
      'Design scalable, secure data architectures on Google Cloud Platform with BigQuery, Cloud Storage and modern data engineering practices.',
  },
  '/services/sovereigncloud': {
    title: 'Sovereign Cloud Solutions | ivmanto.com',
    description:
      'Architectural perspectives on data, operations and AI sovereignty to meet EU compliance and security needs in the cloud.',
  },
  '/services/ml-engineering': {
    title: 'ML Engineering on Vertex AI | ivmanto.com',
    description:
      'Operationalize machine learning on Google Cloud. Automated training, deployment and monitoring pipelines on Vertex AI.',
  },
  '/services/principles': {
    title: 'Guiding Principles | ivmanto.com',
    description:
      'DAMA-aligned principles for data strategy, governance and architecture that make your data a reliable, valuable asset for decision-making and AI.',
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
