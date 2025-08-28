<template>
  <div class="service-section">
    <h1 v-if="service">{{ service.name }}</h1>
    <div v-if="service" v-html="service.content"></div>
    <div v-else>
      <h2>Service Not Found</h2>
      <p>Please select a valid service from the menu.</p>
    </div>
  </div>
</template>

<script setup>
import { computed, toRefs } from 'vue'

const props = defineProps({
  id: {
    type: String,
    required: true,
  },
})

const { id } = toRefs(props)

// In a real application, you would fetch this data from a CMS or API
const allServicesData = {
  'web-development': {
    name: 'Web Development',
    content:
      '<p>We build modern, responsive, and blazing-fast websites tailored to your business needs. Our technology stack includes the latest frameworks and best practices to ensure scalability and performance.</p><ul><li>Custom Web Applications</li><li>E-commerce Solutions</li><li>Content Management Systems</li></ul>',
  },
  'mobile-apps': {
    name: 'Mobile App Development',
    content:
      '<p>From iOS to Android, we create beautiful and functional mobile applications that provide a seamless user experience. We handle the entire lifecycle from concept to deployment.</p><ul><li>Native iOS & Android Apps</li><li>Cross-platform Development</li><li>App Store Submission</li></ul>',
  },
  'ui-ux-design': {
    name: 'UI/UX Design',
    content:
      '<p>Our design team focuses on creating intuitive, engaging, and aesthetically pleasing user interfaces and experiences. We believe that great design is key to user adoption and satisfaction.</p><ul><li>Wireframing & Prototyping</li><li>User Research</li><li>Interaction Design</li></ul>',
  },
  'devops-cloud': {
    name: 'DevOps & Cloud',
    content:
      '<p>We help you streamline your development and deployment processes with modern DevOps practices and cloud infrastructure solutions, ensuring reliability and efficiency.</p><ul><li>Continuous Integration & Deployment (CI/CD)</li><li>Cloud Architecture (AWS, Azure, GCP)</li><li>Infrastructure as Code (IaC)</li></ul>',
  },
}

const service = computed(() => allServicesData[id.value])
</script>

<style scoped>
.service-section {
  animation: fadeIn 0.5s ease-in-out;
}

h1 {
  color: #212529;
}

/* Using :deep to style v-html content */
:deep(p) {
  line-height: 1.6;
  color: #495057;
}

:deep(ul) {
  list-style-position: inside;
  padding-left: 0;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
