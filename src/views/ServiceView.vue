<template>
  <div class="service-view">
    <aside class="left-column">
      <h2>Our Services</h2>
      <nav>
        <ul>
          <li v-for="service in services" :key="service.id">
            <router-link :to="{ name: 'service-detail', params: { id: service.id } }">
              {{ service.name }}
            </router-link>
          </li>
        </ul>
      </nav>
    </aside>

    <main class="main-content">
      <!-- The content for the selected service will be rendered here via nested routes -->
      <router-view :key="$route.path"></router-view>
    </main>

    <aside class="right-column">
      <h3>Related Information</h3>
      <div v-if="currentService" class="related-content">
        <h4>About {{ currentService.name }}</h4>
        <p>
          Here you can find supplementary materials, links, or actions related to
          {{ currentService.name.toLowerCase() }}.
        </p>
        <ul>
          <li><a href="#">Case Studies</a></li>
          <li><a href="#">Technical Docs</a></li>
          <li><a href="#">Contact an Expert</a></li>
        </ul>
      </div>
      <div v-else class="related-content-placeholder">
        <p>Please select a service from the list to see more information.</p>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'

// In a real application, this data would likely come from an API
const services = ref([
  { id: 'web-development', name: 'Web Development' },
  { id: 'mobile-apps', name: 'Mobile App Development' },
  { id: 'ui-ux-design', name: 'UI/UX Design' },
  { id: 'devops-cloud', name: 'DevOps & Cloud' },
])

const route = useRoute()

// This computed property will react to changes in the route
const currentService = computed(() => {
  return services.value.find((s) => s.id === route.params.id)
})
</script>

<style scoped>
.service-view {
  display: flex;
  height: calc(100vh - 80px); /* Example height, adjust based on your layout */
  background-color: #fff;
}

.left-column {
  width: 25%;
  padding: 1.5rem;
  background-color: #f8f9fa;
  border-right: 1px solid #dee2e6;
  overflow-y: auto;
}

.main-content {
  width: 50%;
  padding: 1.5rem 2.5rem;
  overflow-y: auto;
}

.right-column {
  width: 25%;
  padding: 1.5rem;
  background-color: #f8f9fa;
  border-left: 1px solid #dee2e6;
  overflow-y: auto;
}

h2,
h3 {
  margin-top: 0;
  color: #343a40;
  border-bottom: 2px solid #007bff;
  padding-bottom: 0.5rem;
  margin-bottom: 1rem;
}

ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.left-column li a {
  display: block;
  padding: 0.75rem 1rem;
  text-decoration: none;
  color: #495057;
  border-radius: 0.25rem;
  transition:
    background-color 0.2s ease-in-out,
    color 0.2s ease-in-out;
}

.left-column li a:hover {
  background-color: #e9ecef;
  color: #0056b3;
}

.left-column li a.router-link-exact-active {
  background-color: #007bff;
  color: white;
  font-weight: bold;
}

.related-content ul {
  list-style-type: disc;
  padding-left: 20px;
}

.related-content li {
  margin-bottom: 0.5rem;
}
</style>
