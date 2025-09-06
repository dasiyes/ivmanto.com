import { defineAsyncComponent, type Component } from 'vue'

export type Service = {
  id: string
  menuTitle: string
  title: string
  summary: string
  icon: string // SVG path data
  detailsComponent: Component
  tagDetails?: { [key: string]: string }
  industries: string[]
}

// [!] Example Service Definition:
// {
//   id: '',
//   menuTitle: '',
//   title: '',
//   summary: '',
//   icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />`,
//   detailsComponent: defineAsyncComponent(
//     () => import('@/components/services-content/DataArchitecture.vue'),
//   ),
//   tagDetails: {},
//   industries: ['Finance', 'Healthcare', 'Public sector'],
// },

export const services: Service[] = [
  {
    id: 'principles',
    menuTitle: 'Guiding Principles',
    title: 'Our Guiding Principles for Data Management & Architecture',
    summary:
      'Our approach is grounded in the globally recognized standards of DAMA and its Data Management Body of Knowledge (DMBOK), ensuring we turn your information into your most reliable and valuable asset.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10l-2 1m0 0l-2-1m2 1v2.5M20 7l-2 1m2-1l-2-1m2 1v2.5M14 4l-2-1-2 1M4 7l2 1M4 7l2-1M4 7v2.5M12 21l-2-1m2 1l-2 1m2-1v-2.5M6 18l-2-1m2 1l-2 1m2-1V15M2 4h20M2 11h20M2 18h20" />`,
    detailsComponent: defineAsyncComponent(
      () => import('@/components/services-content/Principles.vue'),
    ),
    tagDetails: {
      DMBOOK:
        'The DAMA-DMBOK (Data Management Body of Knowledge) is a framework of data management best practices, often used as a study guide for data management certification.',
    },
    industries: ['All'],
  },
  {
    id: 'sovereigncloud',
    menuTitle: 'Sovereign Cloud',
    title:
      'Sovereignty Cloud in Germany: Navigating Control, Compliance, and Innovation in the Cloud Era',
    summary: 'An Architectural Perspective on Data, Operations, and AI Sovereignty',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />`,
    detailsComponent: defineAsyncComponent(
      () => import('@/components/services-content/SovereignCloudDE.vue'),
    ),
    tagDetails: {
      Cloud:
        "Build what\’s next. Better software. Faster. 1) Use Google\'s core infrastructure, data analytics, and machine learning. 2) Protect your data and apps with the same security technology Google uses. 3) Avoid vendor lock-in and run your apps on open source solutions",
      DataAct:
        'The Data Act is a comprehensive initiative to address the challenges and unleash the opportunities presented by data in the European Union, emphasising fair access and user rights, while ensuring the protection of personal data.',
    },
    industries: ['Finance', 'Healthcare', 'Public sector'],
  },
  {
    id: 'data-architecture',
    menuTitle: 'Data Architecture',
    title: 'Cloud Data Architecture',
    summary:
      'Beyond the Blueprint: Why Your Data Architecture is the True Engine of Your AI Strategy',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />`,
    detailsComponent: defineAsyncComponent(
      () => import('@/components/services-content/DataArchitecture.vue'),
    ),
    tagDetails: {
      BigQuery:
        "Google's fully-managed, petabyte-scale, and cost-effective analytics data warehouse that lets you run analytics over vast amounts of data in near real time.",
      GCS: 'Google Cloud Storage (GCS) is a unified object storage for developers and enterprises, from live data serving to data analytics/ML to data archiving.',
      CloudSQL:
        'Cloud SQL is a fully-managed database service that makes it easy to set up, maintain, manage, and administer your relational PostgreSQL, MySQL, and SQL Server databases in the cloud.',
    },
    industries: ['Finance', 'Retail', 'Healthcare'],
  },
  {
    id: 'ml-engineering',
    menuTitle: 'ML Engineering & MLOps',
    title: 'ML Engineering & MLOps',
    summary: 'Operationalizing machine learning models from prototype to production.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10l-2 1m0 0l-2-1m2 1v2.5M20 7l-2 1m2-1l-2-1m2 1v2.5M14 4l-2-1-2 1M4 7l2 1M4 7l2-1M4 7v2.5M12 21l-2-1m2 1l-2 1m2-1v-2.5M6 18l-2-1m2 1l-2 1m2-1V15M2 4h20M2 11h20M2 18h20" />`,
    detailsComponent: defineAsyncComponent(
      () => import('@/components/services-content/MlEngineering.vue'),
    ),
    tagDetails: {
      VertexAI:
        'A unified AI platform that helps you build, deploy, and scale ML models faster, with pre-trained and custom tooling within a single platform.',
      CICD: 'Continuous Integration and Continuous Delivery (CI/CD) is a method to frequently deliver apps to customers by introducing automation into the stages of app development.',
    },
    industries: ['Retail', 'Healthcare'],
  },
  {
    id: 'data-governance',
    menuTitle: 'Data Governance',
    title: 'Data Governance & Strategy',
    summary: 'Implementing DAMA-aligned principles for data quality and security.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />`,
    detailsComponent: defineAsyncComponent(
      () => import('@/components/services-content/DataGovernance.vue'),
    ),
    tagDetails: {
      DAMA: 'The DAMA-DMBOK (Data Management Body of Knowledge) is a framework of data management best practices, often used as a study guide for data management certification.',
    },
    industries: ['Finance', 'Healthcare'],
  },
]

const servicesMap = new Map<string, Service>(services.map((service) => [service.id, service]))

export function getServiceById(id: string | undefined): Service | undefined {
  if (!id) return undefined
  return servicesMap.get(id)
}
