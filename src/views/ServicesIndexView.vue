<template>
  <div
    class="w-full md:w-4/5 max-w-7xl bg-white shadow-lg rounded-xl flex flex-col overflow-hidden"
  >
    <!-- Top Filter Bar -->
    <div class="flex-shrink-0 p-3 border-b border-gray-200 flex items-center gap-4">
      <span class="font-semibold text-gray-500 text-sm ml-2 flex items-center gap-2">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-5 h-5"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M9.594 1.576c.338-.462.976-.462 1.314 0l1.435 1.958a1 1 0 00.81.466h2.176c.528 0 .955.46.852.984l-.34 1.702a1 1 0 00.294.869l1.54 1.54a1 1 0 010 1.414l-1.54 1.54a1 1 0 00-.294.869l.34 1.702c.103.524-.324.984-.852-.984h-2.176a1 1 0 00-.81.466l-1.435 1.958c-.338.462-.976.462-1.314 0l-1.435-1.958a1 1 0 00-.81-.466H4.03c-.528 0-.955-.46-.852-.984l.34-1.702a1 1 0 00-.294-.869l-1.54-1.54a1 1 0 010-1.414l1.54-1.54a1 1 0 00.294-.869l-.34-1.702c-.103-.524.324-.984.852-.984h2.176a1 1 0 00.81-.466l1.435-1.958zM12 8.25a3.75 3.75 0 100 7.5 3.75 3.75 0 000-7.5z"
          />
        </svg>
        <span>Industries:</span>
      </span>
      <div class="flex items-center gap-2">
        <button
          v-for="industry in industries"
          :key="industry"
          @click="setActiveIndustry(industry)"
          :class="[activeIndustry === industry ? 'filter-button-active' : 'filter-button']"
        >
          {{ industry }}
        </button>
      </div>
    </div>

    <!-- Main Content Area (3 columns) -->
    <div class="flex flex-grow overflow-hidden min-h-0">
      <!-- Left Column: Service List -->
      <div class="flex-shrink-0 w-[20%] bg-white border-r border-gray-200 overflow-y-auto">
        <div class="p-4">
          <h2 class="text-lg font-bold text-dark-slate mb-4">Services</h2>
          <nav class="space-y-1">
            <a
              v-for="service in filteredServices"
              :key="service.id"
              @click.prevent="selectService(service.id)"
              href="#"
              class="block px-3 py-2 rounded-md text-sm font-medium transition-colors cursor-pointer"
              :class="[
                selectedServiceId === service.id
                  ? 'bg-primary text-white'
                  : 'text-gray-600 hover:bg-gray-100',
              ]"
            >
              {{ service.title }}
            </a>
          </nav>
        </div>
      </div>
      <!-- Middle Column: Patterned Background -->
      <div class="flex-grow pattern-bg">
        <ServiceDetail :service="selectedService" @update-right-column="updateRightColumn" />
      </div>

      <!-- Right Column: Empty -->
      <div class="flex-shrink-0 w-[20%] bg-white border-l border-gray-200 overflow-y-auto">
        <RightColumnContent :content="rightColumnContent" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { services, getServiceById } from '@/data/services'
import ServiceDetail from '@/components/services/ServiceDetail.vue'
import RightColumnContent from '@/components/services/RightColumnContent.vue'

const industries = ['All', 'Finance', 'Healthcare', 'Retail']
const activeIndustry = ref('All')

const filteredServices = computed(() => {
  if (activeIndustry.value === 'All') {
    return services
  }
  return services.filter((service) => service.industries.includes(activeIndustry.value))
})

const rightColumnContent = ref<string | undefined>()
const selectedServiceId = ref<string | undefined>(filteredServices.value[0]?.id)

const selectedService = computed(() => getServiceById(selectedServiceId.value))

function selectService(id: string) {
  selectedServiceId.value = id
}

function setActiveIndustry(industry: string) {
  activeIndustry.value = industry
}

function updateRightColumn(content: string | undefined) {
  rightColumnContent.value = content
}

// Watch for changes in the filtered list and update the selection
watch(filteredServices, (newServices) => {
  const isSelectedVisible = newServices.some((s) => s.id === selectedServiceId.value)

  if (!isSelectedVisible) {
    // If the current selection is no longer in the list, select the first available service
    selectedServiceId.value = newServices[0]?.id
    rightColumnContent.value = undefined // Clear right column when selection changes
  }
})
</script>

<style scoped>
.pattern-bg {
  background-color: #ffffff;
  /* A subtle grid pattern */
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 32 32' width='32' height='32' fill='none' stroke='rgb(243 244 246 / 1)'%3e%3cpath d='M0 .5H31.5V32'/%3e%3c/svg%3e");
}
.filter-button {
  @apply px-3 py-1 text-sm rounded-full bg-white border border-gray-300 text-gray-600 hover:bg-gray-100 hover:border-gray-400 transition-colors;
}
.filter-button-active {
  @apply px-3 py-1 text-sm rounded-full bg-primary border border-primary text-white;
}
</style>
