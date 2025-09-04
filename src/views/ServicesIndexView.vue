<template>
  <div
    class="w-full md:w-4/5 max-w-7xl bg-white shadow-lg rounded-xl flex flex-col overflow-hidden h-full"
  >
    <!-- Top Filter Bar -->
    <div
      class="flex-shrink-0 p-4 border-b border-gray-200 flex flex-col lg:flex-row lg:items-center gap-x-4 gap-y-3"
    >
      <div class="w-full lg:w-[20%] shrink-0 flex items-center gap-2">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-7 h-7"
          aria-hidden="true"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M9.594 1.576c.338-.462.976-.462 1.314 0l1.435 1.958a1 1 0 00.81.466h2.176c.528 0 .955.46.852.984l-.34 1.702a1 1 0 00.294.869l1.54 1.54a1 1 0 010 1.414l-1.54 1.54a1 1 0 00-.294.869l.34 1.702c.103.524-.324.984-.852-.984h-2.176a1 1 0 00-.81.466l-1.435 1.958c-.338.462-.976.462-1.314 0l-1.435-1.958a1 1 0 00-.81-.466H4.03c-.528 0-.955-.46-.852-.984l.34-1.702a1 1 0 00-.294-.869l-1.54-1.54a1 1 0 010-1.414l1.54-1.54a1 1 0 00.294-.869l-.34-1.702c-.103-.524.324-.984.852-.984h2.176a1 1 0 00.81-.466l1.435-1.958zM12 8.25a3.75 3.75 0 100 7.5 3.75 3.75 0 000-7.5z"
          />
        </svg>
        <span class="font-semibold text-gray-500 text-sm">Industries:</span>
      </div>
      <div class="flex items-center gap-2 flex-wrap">
        <button
          v-for="industry in industries"
          :key="industry"
          @click="setActiveIndustry(industry)"
          :class="[activeIndustry === industry ? 'filter-button-active' : 'filter-button']"
        >
          {{ industry }}
        </button>
      </div>

      <!-- Book Consultation CTA -->
      <div class="lg:ml-auto pt-2 lg:pt-0">
        <a
          href="/contact?subject=Consultation"
          class="inline-flex items-center justify-center w-full lg:w-auto px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-primary hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-dark transition-colors"
        >
          Book a Consultation
        </a>
      </div>
    </div>

    <!-- Mobile Menu Toggle -->
    <div class="lg:hidden p-4 border-b border-gray-200">
      <button
        @click="isMobileNavOpen = !isMobileNavOpen"
        class="w-full flex justify-between items-center p-2 bg-gray-50 rounded-md text-left"
      >
        <span class="flex items-center gap-x-3 font-semibold text-dark-slate">
          <svg
            v-if="selectedService"
            class="h-5 w-5 flex-shrink-0 text-gray-500"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            aria-hidden="true"
            v-html="selectedService.icon"
          ></svg>
          <span>{{ selectedService?.menuTitle || 'Select a Service' }}</span>
        </span>
        <svg
          class="w-5 h-5 text-gray-600 transform transition-transform"
          :class="{ 'rotate-180': isMobileNavOpen }"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M19 9l-7 7-7-7"
          ></path>
        </svg>
      </button>
    </div>

    <!-- Main Content Area (3 columns) -->
    <div class="flex flex-col lg:flex-row flex-grow overflow-hidden min-h-0">
      <!-- Left Column: Service List -->
      <div
        class="flex-shrink-0 w-full lg:w-[20%] bg-white border-r border-gray-200 overflow-y-auto transition-all duration-300"
        :class="isMobileNavOpen ? 'block' : 'hidden lg:block'"
      >
        <div class="p-4">
          <h2 class="text-lg font-bold text-dark-slate mb-2">Services</h2>
          <hr class="border-gray-200 mb-2" />
          <nav class="space-y-1">
            <a
              v-for="service in filteredServices"
              :key="service.id"
              @click.prevent="selectService(service.id)"
              href="#"
              class="flex items-center gap-x-3 px-3 py-2 rounded-md text-sm font-medium transition-colors cursor-pointer"
              :class="[
                selectedServiceId === service.id
                  ? 'bg-primary text-white'
                  : 'text-gray-600 hover:bg-gray-100',
              ]"
            >
              <svg
                class="h-5 w-5 flex-shrink-0"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                aria-hidden="true"
                v-html="service.icon"
              ></svg>
              <span>
                {{ service.menuTitle }}
              </span>
            </a>
          </nav>
        </div>
      </div>
      <!-- Middle Column: Patterned Background -->
      <div class="flex-grow pattern-bg order-first lg:order-none">
        <ServiceDetail :service="selectedService" @update-right-column="updateRightColumn" />
      </div>

      <!-- Right Column: Empty -->
      <div
        class="flex-shrink-0 w-full lg:w-[20%] bg-white border-l border-gray-200 overflow-y-auto"
      >
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

const isMobileNavOpen = ref(false)
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
  isMobileNavOpen.value = false // Close nav on selection
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
