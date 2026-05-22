import { defineAsyncComponent, type Component } from 'vue'

export type Service = {
  id: string
  menuTitle: string
  summary: string
  icon: string // SVG path data
  detailsComponent: Component
  tagDetails?: { [key: string]: string }
  industries: string[]
  relatedBlogSlugs: string[]
  featured?: boolean
  keywords?: string[]
}

export const services: Service[] = [
  {
    id: 'ai-automation-discovery',
    menuTitle: 'AI & Automation Discovery',
    summary:
      'Helping SMBs cut through the hype to identify high-impact, practical AI and automation use-cases — delivered as a clear, ROI-prioritized roadmap.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />`,
    detailsComponent: defineAsyncComponent(
      () => import('~/components/services-content/AiAutomationDiscovery.vue'),
    ),
    industries: ['Retail', 'Professional Services', 'Manufacturing', 'Logistics'],
    relatedBlogSlugs: [],
    featured: true,
    keywords: [
      'AI use-case discovery',
      'business process automation',
      'workflow automation',
      'automation roadmap',
      'ROI prioritization',
      'AI for SMBs',
      'intelligent automation',
    ],
  },
  {
    id: 'data-pipeline-engineering',
    menuTitle: 'Data Pipeline Engineering',
    summary:
      'Designing robust, scalable data pipelines that transform raw information into high-quality, curated datasets your analytics and AI models can trust.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.25 14.25h13.5m-13.5 0a3 3 0 01-3-3m3 3a3 3 0 100 6h13.5a3 3 0 100-6m-16.5-3a3 3 0 013-3h13.5a3 3 0 013 3m-19.5 0a4.5 4.5 0 01.9-2.7L5.737 5.1a3.375 3.375 0 012.7-1.35h7.126c1.062 0 2.062.5 2.7 1.35l2.587 3.45a4.5 4.5 0 01.9 2.7" />`,
    detailsComponent: defineAsyncComponent(
      () => import('~/components/services-content/DataPipelineEngineering.vue'),
    ),
    industries: ['Finance', 'Healthcare', 'Retail', 'Technology'],
    relatedBlogSlugs: ['VisionaryDataArchitecture', 'FromBigDataTo'],
    featured: true,
    keywords: [
      'data pipeline engineering',
      'data pipeline design',
      'data architecture',
      'data integration',
      'ETL pipelines',
      'curated datasets',
      'AI-ready data',
      'scalable data infrastructure',
    ],
  },
  {
    id: 'agentic-ai-solutions',
    menuTitle: 'Agentic AI Solutions',
    summary:
      'Designing and implementing next-generation Agentic AI — autonomous, reasoning agents — alongside your engineering and product teams.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.25 3v1.5M4.5 8.25H3m18 0h-1.5M4.5 12H3m18 0h-1.5m-15 3.75H3m18 0h-1.5M8.25 19.5V21M12 3v1.5m0 15V21m3.75-18v1.5m0 15V21m-9-1.5h10.5a2.25 2.25 0 002.25-2.25V6.75a2.25 2.25 0 00-2.25-2.25H6.75A2.25 2.25 0 004.5 6.75v10.5a2.25 2.25 0 002.25 2.25zm.75-12h9v9h-9v-9z" />`,
    detailsComponent: defineAsyncComponent(
      () => import('~/components/services-content/AgenticAiSolutions.vue'),
    ),
    industries: ['Technology', 'Finance', 'Professional Services'],
    relatedBlogSlugs: ['the-shift-to-agentic-ai-and-autonomous-workflows', 'TheRiseOfSlm'],
    featured: true,
    keywords: [
      'agentic AI',
      'AI agents',
      'multi-agent systems',
      'autonomous agents',
      'agentic AI implementation',
      'intelligent workflows',
      'AI solution architecture',
      'team enablement',
    ],
  },
  {
    id: 'principles',
    menuTitle: 'Guiding Principles',
    summary:
      'Our approach is grounded in the globally recognized standards of DAMA and its Data Management Body of Knowledge (DMBOK), ensuring we turn your information into your most reliable and valuable asset.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10l-2 1m0 0l-2-1m2 1v2.5M20 7l-2 1m2-1l-2-1m2 1v2.5M14 4l-2-1-2 1M4 7l2 1M4 7l2-1M4 7v2.5M12 21l-2-1m2 1l-2 1m2-1v-2.5M6 18l-2-1m2 1l-2 1m2-1V15M2 4h20M2 11h20M2 18h20" />`,
    detailsComponent: defineAsyncComponent(
      () => import('~/components/services-content/Principles.vue'),
    ),
    tagDetails: {
      DMBOOK:
        'The DAMA-DMBOK (Data Management Body of Knowledge) is a framework of data management best practices, often used as a study guide for data management certification.',
    },
    industries: ['All'],
    relatedBlogSlugs: ['NavigatingTheDataFrontier', 'OnDataManagement'],
  },
  {
    id: 'sovereigncloud',
    menuTitle: 'Sovereign Cloud',
    summary: 'An Architectural Perspective on Data, Operations, and AI Sovereignty',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />`,
    detailsComponent: defineAsyncComponent(
      () => import('~/components/services-content/SovereignCloudDE.vue'),
    ),
    tagDetails: {
      Cloud:
        "Build what\'s next. Better software. Faster. 1) Use Google\'s core infrastructure, data analytics, and machine learning. 2) Protect your data and apps with the same security technology Google uses. 3) Avoid vendor lock-in and run your apps on open source solutions",
      DataAct:
        'The Data Act is a comprehensive initiative to address the challenges and unleash the opportunities presented by data in the European Union, emphasising fair access and user rights, while ensuring the protection of personal data.',
    },
    industries: ['Finance', 'Healthcare', 'Public sector'],
    relatedBlogSlugs: ['FromBigDataTo'],
  },
  {
    id: 'data-architecture',
    menuTitle: 'Data Architecture',
    summary:
      'Beyond the Blueprint: Why Your Data Architecture is the True Engine of Your AI Strategy',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />`,
    detailsComponent: defineAsyncComponent(
      () => import('~/components/services-content/DataArchitecture.vue'),
    ),
    tagDetails: {
      BigQuery:
        "Google's fully-managed, petabyte-scale, and cost-effective analytics data warehouse that lets you run analytics over vast amounts of data in near real time.",
      GCS: 'Google Cloud Storage (GCS) is a unified object storage for developers and enterprises, from live data serving to data analytics/ML to data archiving.',
      CloudSQL:
        'Cloud SQL is a fully-managed database service that makes it easy to set up, maintain, manage, and administer your relational PostgreSQL, MySQL, and SQL Server databases in the cloud.',
    },
    industries: ['Finance', 'Retail', 'Healthcare'],
    relatedBlogSlugs: ['VisionaryDataArchitecture'],
  },
  {
    id: 'ml-engineering',
    menuTitle: 'AI & ML Solutions',
    summary:
      'Operationalizing machine learning models and developing your AI transformation strategy from prototype to production.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10l-2 1m0 0l-2-1m2 1v2.5M20 7l-2 1m2-1l-2-1m2 1v2.5M14 4l-2-1-2 1M4 7l2 1M4 7l2-1M4 7v2.5M12 21l-2-1m2 1l-2 1m2-1v-2.5M6 18l-2-1m2 1l-2 1m2-1V15M2 4h20M2 11h20M2 18h20" />`,
    detailsComponent: defineAsyncComponent(
      () => import('~/components/services-content/MlEngineering.vue'),
    ),
    tagDetails: {
      VertexAI:
        'A unified AI platform that helps you build, deploy, and scale ML models faster, with pre-trained and custom tooling within a single platform.',
      CICD: 'Continuous Integration and Continuous Delivery (CI/CD) is a method to frequently deliver apps to customers by introducing automation into the stages of app development.',
    },
    industries: ['Retail', 'Healthcare'],
    relatedBlogSlugs: ['the-shift-to-agentic-ai-and-autonomous-workflows', 'TheRiseOfSlm'],
  },
  {
    id: 'data-strategy-and-governance',
    menuTitle: 'Data Governance',
    summary:
      'Implementing a robust data governance framework and DAMA-aligned principles for data quality and security.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />`,
    detailsComponent: defineAsyncComponent(
      () => import('~/components/services-content/DataGovernance.vue'),
    ),
    tagDetails: {
      DAMA: 'The DAMA-DMBOK (Data Management Body of Knowledge) is a framework of data management best practices, often used as a study guide for data management certification.',
    },
    industries: ['Finance', 'Healthcare'],
    relatedBlogSlugs: ['DataMeshGovernance', 'OnDataManagement'],
  },
]

const servicesMap = new Map<string, Service>(services.map((service) => [service.id, service]))

export function getServiceById(id: string | undefined): Service | undefined {
  if (!id) return undefined
  return servicesMap.get(id)
}
