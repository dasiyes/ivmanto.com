<script setup lang="ts">
interface Service {
  title: string
  description: string
}

const services: Service[] = [
  {
    title: 'Cloud Data Architecture',
    description:
      'Designing and implementing scalable data lakes, warehouses, and lakehouses on GCP using services like BigQuery, GCS, and Dataproc.',
  },
  {
    title: 'Data Engineering & Pipelines',
    description:
      'Building robust and efficient ETL/ELT pipelines with Cloud Data Fusion, Dataflow, and Composer to ensure timely and reliable data.',
  },
  {
    title: 'AI & ML Consulting',
    description:
      'Harnessing the power of Vertex AI and other GCP machine learning services to build, deploy, and manage predictive models.',
  },
  {
    title: 'Business Intelligence & Analytics',
    description:
      'Creating insightful dashboards and reports with Looker and connecting your data sources for a unified view of your business.',
  },
  {
    title: 'Cloud Strategy & Governance',
    description:
      'Providing expert guidance on cloud adoption, cost optimization, security, and data governance best practices.',
  },
  {
    title: 'DAMA Framework Implementation',
    description:
      'Assisting your organization in aligning with DAMA principles for certified data management excellence.',
  },
]
</script>

<template>
  <section id="services" class="py-20 md:py-28">
    <div class="container mx-auto px-6">
      <div class="text-center mb-16">
        <h2 class="text-3xl md:text-4xl font-bold text-dark-slate">What I Can Do For You</h2>
        <p class="text-lg text-gray-600 mt-4 max-w-2xl mx-auto">
          Leveraging deep expertise in Google Cloud to solve your most complex data challenges.
        </p>
      </div>
      <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div v-for="service in services" :key="service.title" class="bg-light-gray p-8 rounded-xl">
          <h3 class="text-2xl font-bold text-dark-slate">{{ service.title }}</h3>
          <p class="mt-4 text-gray-600">{{ service.description }}</p>
        </div>
      </div>
    </div>
  </section>
</template>
