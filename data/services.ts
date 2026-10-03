import { defineAsyncComponent, type Component } from 'vue'

// entry:     low-risk first purchase that leads into a project
// core:      the three hands-on delivery pillars
// ongoing:   recurring support after handoff
// expertise: long-form background pages (principles, governance, sovereignty)
export type ServiceTier = 'entry' | 'core' | 'ongoing' | 'expertise'

export type ServiceDeliverable = { title: string; text: string }
export type ServiceFaq = { question: string; answer: string }
export type ServicePrice = { label: string; amount: number; currency: string; note: string }

export type Service = {
  id: string
  tier: ServiceTier
  menuTitle: string
  summary: string
  icon: string // SVG path data
  seoTitle: string
  seoDescription: string
  // Structured offer content (entry, core, ongoing tiers)
  eyebrow?: string
  headline?: string
  intro?: string
  problems?: string[]
  deliverables?: ServiceDeliverable[]
  useCases?: string[]
  stack?: string[]
  faqs?: ServiceFaq[]
  price?: ServicePrice
  nextStepId?: string
  // Long-form article content (expertise tier)
  detailsComponent?: Component
  tagDetails?: { [key: string]: string }
  industries: string[]
  relatedBlogSlugs: string[]
  keywords?: string[]
}

export const services: Service[] = [
  {
    id: 'ai-readiness-audit',
    tier: 'entry',
    menuTitle: 'Data & AI Readiness Audit',
    summary:
      'A short, fixed-scope audit of your data and workflows that ends with a clear answer: what to automate first, what it takes, and what it pays back.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />`,
    seoTitle: 'Data & AI Readiness Audit for SMBs | ivmanto.com',
    seoDescription:
      'A 1 to 2 day audit of your data and workflows by an independent senior engineer. Find the highest ROI automation and AI agent use case. From €100, free if we sign the project.',
    eyebrow: 'Start here',
    headline: 'Find out what is worth automating before you spend on AI',
    intro:
      'Most AI projects in small and medium businesses stall for the same two reasons: the data underneath is not ready, or the chosen use case never had a business case. In 1 to 2 days I review your workflows, systems and data, and tell you plainly what is feasible, what it costs and where the return is highest.',
    problems: [
      'You know AI could help, but not where to start or what is realistic',
      'A previous chatbot or proof of concept never made it into daily use',
      'Your data sits in a CRM, an ERP, spreadsheets and inboxes that do not talk to each other',
      'You need a number for the budget conversation, not another slide deck',
    ],
    deliverables: [
      {
        title: 'Workflow and bottleneck review',
        text: 'Structured sessions with the people who do the work, mapping where time and money leak today.',
      },
      {
        title: 'Data readiness score',
        text: 'An honest assessment of data quality, access and structure, and what has to be fixed before any agent can rely on it.',
      },
      {
        title: 'Use cases ranked by ROI',
        text: 'Candidate automations and AI agents ranked by effort, risk and payback, with the top one scoped in detail.',
      },
      {
        title: 'Fixed-price proposal',
        text: 'A concrete plan, timeline and price for the first build, so you can decide with real numbers.',
      },
    ],
    useCases: [
      'Deciding between an AI agent, a simple automation or a data cleanup first',
      'Checking whether your data can support a private knowledge assistant',
      'Validating a vendor or agency proposal before you sign it',
    ],
    price: {
      label: 'from €100',
      amount: 100,
      currency: 'EUR',
      note: 'Fully credited when we sign the project, so the audit becomes free.',
    },
    faqs: [
      {
        question: 'What does the audit cost?',
        answer:
          'It starts from €100, depending on the number of systems and workflows in scope. If you go ahead with the project, the full audit fee is credited, so the audit is free.',
      },
      {
        question: 'Do I need a technical team to take part?',
        answer:
          'No. I need access to the people who run the workflows and read access to the relevant systems. I handle the technical review myself.',
      },
      {
        question: 'What if the audit shows AI is not the right answer?',
        answer:
          'Then I will tell you. Sometimes the best return is a data cleanup or a simple automation, and the report says so.',
      },
    ],
    nextStepId: 'cloud-data-engineering',
    industries: ['Retail', 'Professional Services', 'Manufacturing', 'Logistics'],
    relatedBlogSlugs: [],
    keywords: [
      'AI readiness audit',
      'data readiness assessment',
      'AI implementation consultant',
      'automate manual business workflows',
      'AI for small business',
      'automation roadmap',
    ],
  },
  {
    id: 'cloud-data-engineering',
    tier: 'core',
    menuTitle: 'Cloud Data Engineering',
    summary:
      'Clean, consolidated data on Google Cloud: automated pipelines from your CRM, ERP and apps into BigQuery, ready for reporting, automation and AI.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.25 14.25h13.5m-13.5 0a3 3 0 01-3-3m3 3a3 3 0 100 6h13.5a3 3 0 100-6m-16.5-3a3 3 0 013-3h13.5a3 3 0 013 3m-19.5 0a4.5 4.5 0 01.9-2.7L5.737 5.1a3.375 3.375 0 012.7-1.35h7.126c1.062 0 2.062.5 2.7 1.35l2.587 3.45a4.5 4.5 0 01.9 2.7" />`,
    seoTitle: 'Freelance Cloud Data Engineer: BigQuery, dbt & Airflow | ivmanto.com',
    seoDescription:
      'Hands-on cloud data engineering for small and medium businesses. Automated ETL pipelines, BigQuery data warehouses, dbt models and Airflow orchestration on GCP, built in your own cloud.',
    eyebrow: 'Pillar 01',
    headline: 'Clean, connected data that makes automation and AI actually work',
    intro:
      'AI agents and dashboards are only as reliable as the data feeding them. When records are scattered across a CRM, an ERP, spreadsheets and legacy databases, I engineer a single, trustworthy source of truth on Google Cloud, with pipelines that run on schedule and tell you when something breaks.',
    problems: [
      'Reports disagree because every team pulls numbers from a different system',
      'Someone spends hours each week copying data between tools',
      'An AI pilot gave wrong answers because the underlying data was messy',
      'Pipelines fail silently and nobody notices until a report looks wrong',
    ],
    deliverables: [
      {
        title: 'Automated ingestion and ETL/ELT',
        text: 'Scheduled pipelines from third-party APIs, webhooks and operational databases into central storage, orchestrated with Airflow.',
      },
      {
        title: 'BigQuery data warehouse',
        text: 'A performant, cost-controlled warehouse designed around how your business actually reports and decides.',
      },
      {
        title: 'Data modeling with dbt',
        text: 'Cleaned, standardized and tested models that analysts, dashboards and AI agents can all consume.',
      },
      {
        title: 'Monitoring and alerting',
        text: 'Infrastructure as code, automated deployments and alerts, so failures are caught before your users see them.',
      },
    ],
    useCases: [
      'One revenue and customer view across CRM, webshop and accounting',
      'Automated data entry from emails and forms into a database',
      'Preparing company data so an AI agent or knowledge assistant can use it safely',
    ],
    stack: ['Google Cloud', 'BigQuery', 'Cloud Run', 'Cloud SQL', 'dbt', 'Airflow', 'Python', 'Go'],
    faqs: [
      {
        question: 'Do we need to move everything to Google Cloud?',
        answer:
          'No. Source systems stay where they are. I connect them and consolidate the data you need for reporting and AI into BigQuery.',
      },
      {
        question: 'Who owns the pipelines and the warehouse?',
        answer:
          'You do. Everything is built in your own Google Cloud project, with the code in your repository and full documentation at handoff.',
      },
    ],
    nextStepId: 'custom-ai-agents',
    tagDetails: {
      BigQuery:
        "Google's fully managed, petabyte-scale and cost-effective analytics data warehouse for running analytics over large amounts of data in near real time.",
      dbt: 'A transformation framework that turns raw warehouse tables into tested, documented data models using SQL.',
      Airflow:
        'An open source workflow orchestrator for scheduling, monitoring and retrying data pipelines.',
    },
    industries: ['Retail', 'Finance', 'Manufacturing', 'Logistics'],
    relatedBlogSlugs: ['VisionaryDataArchitecture', 'FromBigDataTo'],
    keywords: [
      'freelance cloud data architect',
      'contract data engineer for small business',
      'BigQuery data pipelines',
      'automated ETL',
      'dbt data modeling',
      'Airflow orchestration',
      'data infrastructure for SMBs',
    ],
  },
  {
    id: 'custom-ai-agents',
    tier: 'core',
    menuTitle: 'Custom AI Agents & Workflow Automation',
    summary:
      'Task-focused AI agents that take real action in your tools: reading documents, updating your CRM, handling intake and support, with guardrails and human approval where it matters.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.25 3v1.5M4.5 8.25H3m18 0h-1.5M4.5 12H3m18 0h-1.5m-15 3.75H3m18 0h-1.5M8.25 19.5V21M12 3v1.5m0 15V21m3.75-18v1.5m0 15V21m-9-1.5h10.5a2.25 2.25 0 002.25-2.25V6.75a2.25 2.25 0 00-2.25-2.25H6.75A2.25 2.25 0 004.5 6.75v10.5a2.25 2.25 0 002.25 2.25zm.75-12h9v9h-9v-9z" />`,
    seoTitle: 'Hire a Freelance AI Agent Developer | ivmanto.com',
    seoDescription:
      'Custom AI agents and workflow automation for small and medium businesses. Tool-calling agents connected to your CRM, inbox and databases, with guardrails, testing and human approval built in.',
    eyebrow: 'Pillar 02',
    headline: 'AI agents built for real work, not just chat',
    intro:
      'I design and build agents that do a defined job inside your existing tools: read incoming documents, look things up, update records and hand off to a person when a decision is sensitive. Each agent gets narrow permissions, a test suite and logging, so you can see exactly what it did and why.',
    problems: [
      'Your team retypes the same information from emails, PDFs and forms every day',
      'Leads and support requests wait hours for a first response',
      'You tried a generic chatbot, but it could only talk, not act',
      'You worry an AI system will send the wrong price or email the wrong customer',
    ],
    deliverables: [
      {
        title: 'Tool-connected agents',
        text: 'Agents that call your APIs, CRM, booking system, ticketing tool or database to complete a task end to end.',
      },
      {
        title: 'Document processing and extraction',
        text: 'Invoices, orders, intake forms and unstructured emails parsed into structured records and validated before they are written.',
      },
      {
        title: 'Guardrails and human approval',
        text: 'Narrow permissions per agent and approval checkpoints for sensitive actions such as refunds, quotes or contract changes.',
      },
      {
        title: 'Evaluation and monitoring',
        text: 'Test sets built from your real cases to measure accuracy, catch regressions and reduce hallucinations before release.',
      },
    ],
    useCases: [
      'Lead qualification and intake agent that updates your CRM',
      'Customer support and booking agent connected to your calendar',
      'Supplier invoice and order processing into your database or ERP',
      'Internal operations agent that prepares reports and drafts replies for approval',
    ],
    stack: [
      'Vertex AI',
      'Gemini',
      'Claude Code',
      'Hermes Agent',
      'Cloud Run',
      'BigQuery',
      'Go',
      'Python',
    ],
    faqs: [
      {
        question: 'How do you stop an agent from making costly mistakes?',
        answer:
          'Each agent can only use the tools and data it needs for its job. Sensitive actions wait for human approval, every action is logged, and the agent is tested against your real cases before it goes live.',
      },
      {
        question: 'Are we locked into a platform or monthly license?',
        answer:
          'No. The agents run in your own cloud account, the code is yours, and you pay the model provider directly at cost.',
      },
      {
        question: 'Can the agent work with our existing software?',
        answer:
          'Usually yes. If a system has an API, a database or even a structured export, an agent can work with it. The readiness audit confirms this up front.',
      },
    ],
    nextStepId: 'private-ai-knowledge-base',
    tagDetails: {
      VertexAI:
        "Google Cloud's AI platform for running Gemini and other models inside your own cloud project, with EU region options.",
      'Hermes Agent':
        'An open source agent framework for building tool-using agents that can run on your own infrastructure.',
    },
    industries: ['Professional Services', 'Retail', 'Logistics', 'Technology'],
    relatedBlogSlugs: ['the-shift-to-agentic-ai-and-autonomous-workflows', 'TheRiseOfSlm'],
    keywords: [
      'hire custom AI agent developer',
      'freelance AI automation consultant',
      'AI agent for customer support',
      'lead qualification AI agent',
      'automated document processing',
      'tool-calling AI agents',
      'workflow automation with AI',
    ],
  },
  {
    id: 'private-ai-knowledge-base',
    tier: 'core',
    menuTitle: 'Private AI Knowledge Base',
    summary:
      'A private assistant that answers questions from your own documents with sources, hosted in your cloud in the EU, so company data never trains public models.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.042A8.967 8.967 0 006 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 016 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 016-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0018 18a8.967 8.967 0 00-6 2.292m0-14.25v14.25" />`,
    seoTitle: 'Private RAG & Internal Knowledge Base AI, EU Hosted | ivmanto.com',
    seoDescription:
      'Secure RAG systems for small and medium businesses. A private AI knowledge base over your contracts, manuals and tickets, with cited answers and access control, hosted in your own EU cloud.',
    eyebrow: 'Pillar 03',
    headline: 'Your company knowledge, answered with sources and kept private',
    intro:
      'Your team loses hours searching contracts, manuals, policies and old tickets, or asking the one colleague who knows. I build retrieval-augmented generation (RAG) systems that answer from your own documents only, show the source for every answer and respect who is allowed to see what. Everything runs in your own cloud, in an EU region by default.',
    problems: [
      'The same internal questions get asked and answered again every week',
      'Knowledge leaves the company when an experienced employee does',
      'Staff paste confidential documents into public AI tools',
      'GDPR and client contracts require data to stay in the EU',
    ],
    deliverables: [
      {
        title: 'Retrieval pipeline',
        text: 'Document parsing, chunking and embeddings tuned to your formats, combining keyword and semantic search to find exact names, numbers and clauses.',
      },
      {
        title: 'Cited, grounded answers',
        text: 'Answers built only from retrieved passages, with links to the source, and a clear "I do not know" when the documents do not cover it.',
      },
      {
        title: 'Access control',
        text: 'Document-level permissions so each person only gets answers from what they are allowed to read.',
      },
      {
        title: 'Private EU deployment',
        text: 'Hosted in your own Google Cloud project in an EU region. Your data is not used to train public models.',
      },
    ],
    useCases: [
      'Support team assistant over product manuals and past tickets',
      'Policy and contract Q&A for operations, HR or legal',
      'Onboarding assistant that answers new hires from internal documentation',
    ],
    stack: ['Vertex AI', 'Gemini', 'BigQuery', 'Cloud Run', 'Cloud SQL', 'Go', 'Python'],
    faqs: [
      {
        question: 'Will our documents be used to train AI models?',
        answer:
          'No. The system runs in your own cloud project using enterprise model endpoints that do not train on your data.',
      },
      {
        question: 'Can the data stay in Germany or the EU?',
        answer:
          'Yes. I deploy in EU regions by default, including Frankfurt, and can align the setup with your GDPR and data sovereignty requirements.',
      },
      {
        question: 'How do you reduce wrong answers?',
        answer:
          'The assistant answers only from retrieved passages, shows its sources and is evaluated against a set of real questions from your team before launch.',
      },
    ],
    nextStepId: 'support-and-evolution',
    tagDetails: {
      RAG: 'Retrieval-augmented generation: the model first retrieves relevant passages from your documents and then answers using only those passages.',
      'Data Sovereignty':
        'Keeping data, operations and AI processing under your control and inside the jurisdiction you choose, such as the EU.',
    },
    industries: ['Professional Services', 'Finance', 'Healthcare', 'Manufacturing'],
    relatedBlogSlugs: ['FromBigDataTo'],
    keywords: [
      'internal knowledge base AI',
      'secure RAG pipeline',
      'private AI chatbot with company data',
      'semantic search',
      'private document Q&A',
      'GDPR compliant AI',
      'EU hosted AI',
    ],
  },
  {
    id: 'support-and-evolution',
    tier: 'ongoing',
    menuTitle: 'Support & Evolution Retainer',
    summary:
      'Monthly hands-on support after launch: monitoring, fixes, model updates and new automations, from the engineer who built your system.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0l3.181 3.183a8.25 8.25 0 0013.803-3.7M4.031 9.865a8.25 8.25 0 0113.803-3.7l3.181 3.182m0-4.991v4.99" />`,
    seoTitle: 'Fractional Data & AI Engineer Retainer | ivmanto.com',
    seoDescription:
      'A monthly retainer with a senior data and AI engineer. Monitoring, maintenance, model updates and new automations for the pipelines and agents running your business.',
    eyebrow: 'Ongoing',
    headline: 'A senior data and AI engineer on call, without hiring one',
    intro:
      'Data pipelines and AI agents need care after launch: sources change, models improve and new use cases appear. With a monthly retainer you keep direct access to the engineer who knows your system, for a fraction of a full-time hire.',
    problems: [
      'You cannot justify a full-time data or AI engineer yet',
      'Your system was built by someone who is no longer available',
      'New automation ideas pile up with nobody to build them',
    ],
    deliverables: [
      {
        title: 'Monitoring and maintenance',
        text: 'Pipeline and agent health checks, fixes when sources or APIs change, and cost reviews.',
      },
      {
        title: 'Model and quality updates',
        text: 'Moving to better or cheaper models when it pays off, re-running evaluations to confirm quality holds.',
      },
      {
        title: 'Continuous improvements',
        text: 'A fixed monthly budget of hours for new automations, reports and agent capabilities.',
      },
      {
        title: 'Direct access',
        text: 'A named senior engineer you can message, not a ticket queue.',
      },
    ],
    stack: ['Google Cloud', 'BigQuery', 'Vertex AI', 'dbt', 'Airflow', 'Claude Code'],
    faqs: [
      {
        question: 'Do I need a retainer after a project?',
        answer:
          'No. Every project ends with a full handoff and documentation, so your team can run it alone. The retainer is for teams who prefer to keep an expert involved.',
      },
    ],
    industries: ['All'],
    relatedBlogSlugs: [],
    keywords: [
      'fractional AI engineer',
      'fractional data engineer',
      'AI system maintenance',
      'data pipeline support retainer',
    ],
  },
  {
    id: 'principles',
    tier: 'expertise',
    menuTitle: 'Guiding Principles',
    summary:
      'How I work: grounded in the DAMA Data Management Body of Knowledge (DMBOK), so your information becomes a reliable and valuable asset.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10l-2 1m0 0l-2-1m2 1v2.5M20 7l-2 1m2-1l-2-1m2 1v2.5M14 4l-2-1-2 1M4 7l2 1M4 7l2-1M4 7v2.5M12 21l-2-1m2 1l-2 1m2-1v-2.5M6 18l-2-1m2 1l-2 1m2-1V15M2 4h20M2 11h20M2 18h20" />`,
    seoTitle: 'Guiding Principles | ivmanto.com',
    seoDescription:
      'DAMA-aligned principles for data strategy, governance and architecture that make your data a reliable, valuable asset for decision-making and AI.',
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
    tier: 'expertise',
    menuTitle: 'Sovereign Cloud',
    summary: 'An architectural perspective on data, operations and AI sovereignty in the EU.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />`,
    seoTitle: 'Sovereign Cloud Solutions | ivmanto.com',
    seoDescription:
      'Architectural perspectives on data, operations and AI sovereignty to meet EU compliance and security needs in the cloud.',
    detailsComponent: defineAsyncComponent(
      () => import('~/components/services-content/SovereignCloudDE.vue'),
    ),
    tagDetails: {
      Cloud:
        "Build what's next. Better software. Faster. 1) Use Google's core infrastructure, data analytics, and machine learning. 2) Protect your data and apps with the same security technology Google uses. 3) Avoid vendor lock-in and run your apps on open source solutions",
      DataAct:
        'The Data Act is a comprehensive initiative to address the challenges and unleash the opportunities presented by data in the European Union, emphasising fair access and user rights, while ensuring the protection of personal data.',
    },
    industries: ['Finance', 'Healthcare', 'Public sector'],
    relatedBlogSlugs: ['FromBigDataTo'],
  },
  {
    id: 'data-architecture',
    tier: 'expertise',
    menuTitle: 'Data Architecture',
    summary: 'Why your data architecture is the true engine of your AI strategy.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />`,
    seoTitle: 'Data Architecture on GCP | ivmanto.com',
    seoDescription:
      'Design scalable, secure data architectures on Google Cloud Platform with BigQuery, Cloud Storage and modern data engineering practices.',
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
    tier: 'expertise',
    menuTitle: 'ML Engineering',
    summary: 'Operationalizing machine learning models on Vertex AI, from prototype to production.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10l-2 1m0 0l-2-1m2 1v2.5M20 7l-2 1m2-1l-2-1m2 1v2.5M14 4l-2-1-2 1M4 7l2 1M4 7l2-1M4 7v2.5M12 21l-2-1m2 1l-2 1m2-1v-2.5M6 18l-2-1m2 1l-2 1m2-1V15M2 4h20M2 11h20M2 18h20" />`,
    seoTitle: 'ML Engineering on Vertex AI | ivmanto.com',
    seoDescription:
      'Operationalize machine learning on Google Cloud. Automated training, deployment and monitoring pipelines on Vertex AI.',
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
    tier: 'expertise',
    menuTitle: 'Data Governance',
    summary:
      'A practical data governance framework and DAMA-aligned principles for data quality and security.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />`,
    seoTitle: 'Data Strategy & Governance | ivmanto.com',
    seoDescription:
      'A clear data strategy and practical governance framework that aligns your data initiatives with business goals, compliance and AI readiness.',
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

// Shared stack shown on the services index
export const coreStack = [
  'Google Cloud',
  'BigQuery',
  'Cloud Run',
  'Vertex AI',
  'Gemini',
  'Claude Code',
  'Hermes Agent',
  'dbt',
  'Airflow',
  'Go',
  'Python',
]

// Offers sold as products, in ladder order
export const offerServices = services.filter((s) => s.tier !== 'expertise')
export const expertiseServices = services.filter((s) => s.tier === 'expertise')

const servicesMap = new Map<string, Service>(services.map((service) => [service.id, service]))

export function getServiceById(id: string | undefined): Service | undefined {
  if (!id) return undefined
  return servicesMap.get(id)
}
