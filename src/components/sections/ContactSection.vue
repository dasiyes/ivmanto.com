<script setup lang="ts">
// Ensure you have lang="ts"
import { ref, onMounted } from 'vue'
import { useGemini } from '@/composables/useGemini'

// Fix #3 Part 1: Define an interface for a chat message
interface Message {
  id: number
  sender: 'ai' | 'user' // Use a union type for specific values
  text: string
}

const { isLoading, callGemini } = useGemini()

// Fix #3 Part 2: Tell the ref it will hold an array of Message objects
const messages = ref<Message[]>([])
const userInput = ref('')
const draftedMessage = ref('')
const showFinalForm = ref(false)
const inquirySent = ref(false)

// Fix #3 Part 3: Add types to the function parameters
function addMessage(sender: 'ai' | 'user', text: string) {
  messages.value.push({ sender, text, id: Date.now() })
}

async function handleUserInput() {
  const userText = userInput.value.trim()
  if (!userText || isLoading.value) return

  addMessage('user', userText)
  userInput.value = ''

  const prompt = `A potential client provided the following project description: "${userText}". Draft a concise and professional inquiry message to a freelance data consultant based on this. The message should clearly state the client's needs and ask about the consultant's availability or services. Start the message with "Hello, I'm interested in your services..."`

  try {
    const response = await callGemini(prompt)
    draftedMessage.value = response
    showFinalForm.value = true
    addMessage('ai', 'Here is the draft. You can edit it below before sending.')
  } catch (e) {
    addMessage(
      'ai',
      `"I'm sorry, I couldn't draft a message. Please try describing your project again. Error:"${e}`,
    )
  }
}

function sendInquiry() {
  console.log('Sending inquiry:', draftedMessage.value)
  inquirySent.value = true
}

onMounted(() => {
  addMessage(
    'ai',
    "Hello! To help me understand your needs, please briefly describe the project or challenge you're facing.",
  )
})
</script>

<template>
  <section id="contact" class="py-20 md:py-28 bg-dark-slate text-white">
    <div class="container mx-auto px-6">
      <div class="text-center mb-12">
        <h2 class="text-3xl md:text-4xl font-bold">Let's Build Something Great Together</h2>
        <p class="text-lg text-gray-300 mt-4 max-w-2xl mx-auto">
          Have a project in mind? Let my AI assistant help you draft your inquiry.
        </p>
      </div>

      <div class="max-w-xl mx-auto bg-white/10 p-6 md:p-8 rounded-xl backdrop-blur-sm">
        <div v-if="!inquirySent">
          <div class="h-80 overflow-y-auto pr-2 space-y-4 mb-4">
            <div
              v-for="message in messages"
              :key="message.id"
              class="flex"
              :class="message.sender === 'user' ? 'justify-end' : 'justify-start'"
            >
              <div
                class="p-3 rounded-lg max-w-xs"
                :class="
                  message.sender === 'user' ? 'bg-primary text-white' : 'bg-secondary text-white'
                "
              >
                {{ message.text }}
              </div>
            </div>
          </div>

          <div v-if="!showFinalForm" class="flex gap-4">
            <input
              v-model="userInput"
              @keyup.enter="handleUserInput"
              :disabled="isLoading"
              type="text"
              placeholder="Type your response..."
              class="flex-grow bg-white/20 border-transparent rounded-md py-2 px-3 text-white focus:ring-accent focus:border-accent"
            />
            <button
              @click="handleUserInput"
              :disabled="isLoading"
              class="bg-accent text-white font-bold py-2 px-5 rounded-lg hover:bg-opacity-90 transition-all flex items-center justify-center w-24"
            >
              <span v-if="!isLoading">Send</span>
              <div v-else class="loader border-t-white w-5 h-5"></div>
            </button>
          </div>

          <div v-else>
            <label for="final-message" class="block text-sm font-medium text-gray-300 mb-2"
              >Here is the drafted message. Feel free to edit it before sending.</label
            >
            <textarea
              v-model="draftedMessage"
              rows="6"
              class="w-full bg-white/20 border-transparent rounded-md py-2 px-3 text-white focus:ring-accent focus:border-accent"
            ></textarea>
            <button
              @click="sendInquiry"
              class="w-full mt-4 bg-accent text-white font-bold py-3 px-6 rounded-lg hover:bg-opacity-90 transition-all text-lg"
            >
              Send Inquiry
            </button>
          </div>
        </div>

        <div v-else class="text-center p-8">
          <h3 class="text-2xl font-bold text-white">Thank You!</h3>
          <p class="text-gray-300 mt-2">
            Your message has been sent. I will get back to you shortly.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>
