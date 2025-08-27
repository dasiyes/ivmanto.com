<template>
  <section class="py-20 md:py-28 bg-white">
    <div class="container mx-auto px-6">
      <div class="text-center max-w-3xl mx-auto">
        <h2 class="text-3xl md:text-4xl font-bold text-dark-slate">Core Expertise</h2>
        <p class="text-lg text-gray-600 mt-4">
          I leverage a modern data stack to design and implement robust, scalable, and efficient
          data solutions that drive business value.
        </p>
      </div>

      <div class="mt-16 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <!-- Data Warehousing & Lakes -->
        <div class="bg-light-gray p-8 rounded-xl shadow-sm hover:shadow-lg transition-shadow">
          <h3 class="text-xl font-bold text-primary mb-4">Data Warehousing & Lakes</h3>
          <p class="text-gray-600">
            Building centralized repositories for structured and unstructured data using platforms
            like BigQuery, Snowflake, and data lake architectures on GCS/S3.
          </p>
        </div>

        <!-- Data Processing & ETL/ELT -->
        <div class="bg-light-gray p-8 rounded-xl shadow-sm hover:shadow-lg transition-shadow">
          <h3 class="text-xl font-bold text-primary mb-4">Data Processing & ETL/ELT</h3>
          <p class="text-gray-600">
            Developing scalable data pipelines for batch and stream processing with tools like
            Dataflow (Apache Beam), Dataproc (Spark), and dbt.
          </p>
        </div>

        <!-- Infrastructure & DataOps -->
        <div class="bg-light-gray p-8 rounded-xl shadow-sm hover:shadow-lg transition-shadow">
          <h3 class="text-xl font-bold text-primary mb-4">Infrastructure & DataOps</h3>
          <p class="text-gray-600">
            Automating the deployment and management of data infrastructure using Terraform (IaC)
            and establishing CI/CD practices for data workflows.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>
