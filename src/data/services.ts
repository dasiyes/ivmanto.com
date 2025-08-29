export type Service = {
  id: string
  title: string
  summary: string
  icon: string // SVG path data
  details: string // Placeholder for more detailed content
  tagDetails?: { [key: string]: string }
  industries: string[]
}

export const services: Service[] = [
  {
    id: 'data-architecture',
    title: 'Cloud Data Architecture',
    summary: 'Designing scalable, secure, and cost-effective data platforms on GCP.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l4-4 4 4m0 6l-4 4-4-4" />`,
    details:
      'We specialize in building robust data architectures using Google Cloud services like #BigQuery, #GCS, and #CloudSQL. Our designs prioritize performance, security, and cost-efficiency to provide a solid foundation for your data initiatives.',
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
    title: 'ML Engineering & MLOps',
    summary: 'Operationalizing machine learning models from prototype to production.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10l-2 1m0 0l-2-1m2 1v2.5M20 7l-2 1m2-1l-2-1m2 1v2.5M14 4l-2-1-2 1M4 7l2 1M4 7l2-1M4 7v2.5M12 21l-2-1m2 1l-2 1m2-1v-2.5M6 18l-2-1m2 1l-2 1m2-1V15M2 4h20M2 11h20M2 18h20" />`,
    details:
      'Our MLOps services streamline the machine learning lifecycle. We use #VertexAI to build automated #CICD pipelines for model training, deployment, and monitoring, ensuring your models deliver continuous value.',
    tagDetails: {
      VertexAI:
        'A unified AI platform that helps you build, deploy, and scale ML models faster, with pre-trained and custom tooling within a single platform.',
      CICD: 'Continuous Integration and Continuous Delivery (CI/CD) is a method to frequently deliver apps to customers by introducing automation into the stages of app development.',
    },
    industries: ['Retail', 'Healthcare'],
  },
  {
    id: 'data-governance',
    title: 'Data Governance & Strategy',
    summary: 'Implementing DAMA-aligned principles for data quality and security.',
    icon: `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />`,
    details:
      'We help you establish a strong data governance framework based on #DAMA principles. This ensures your data is accurate, consistent, and secure, transforming it into a trustworthy asset for decision-making.',
    tagDetails: {
      DAMA: 'The DAMA-DMBOK (Data Management Body of Knowledge) is a framework of data management best practices, often used as a study guide for data management certification.',
    },
    industries: ['Finance', 'Healthcare'],
  },
]

export function getServiceById(id: string | undefined): Service | undefined {
  if (!id) return undefined
  return services.find((service) => service.id === id)
}
