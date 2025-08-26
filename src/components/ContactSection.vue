<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue'
import { callGeminiAPI } from '../services/gemini'

interface ChatMessage {
  sender: 'user' | 'ai'
  text: string
}

const chatHistory = ref<ChatMessage[]>([])
const userInput = ref<string>('')
const isLoading = ref<boolean>(false)
const conversationState = ref<'ASK_DETAILS' | 'SHOW_DRAFT' | 'SENT'>('ASK_DETAILS')
const finalMessage = ref<string>('')
const chatWindow = ref<HTMLElement | null>(null)

const scrollToBottom = async (): Promise<void> => {
  await nextTick()
  if (chatWindow.value) {
    chatWindow.value.scrollTop = chatWindow.value.scrollHeight
  }
}

const addMessage = (sender: 'user' | 'ai', text: string): void => {
  chatHistory.value.push({ sender, text })
  scrollToBottom()
}

const handleUserInput = async (): Promise<void> => {
  const message = userInput.value.trim()
  if (!message || isLoading.value) return

  addMessage('user', message)
  userInput.value = ''
  isLoading.value = true

  if (conversationState.value === 'ASK_DETAILS') {
    addMessage(
      'ai',
      'Thank you. Based on your description, I will now draft a professional inquiry for you. Please wait a moment...',
    )

    const prompt = `A potential client provided the following project description: "${message}". 
        Draft a concise and professional inquiry message to a freelance data consultant based on this. 
        The message should clearly state the client's needs and ask about the consultant's availability or services.
        Start the message with "Hello, I'm interested in your services..."`

    try {
      const draftedMessage = await callGeminiAPI(prompt)
      finalMessage.value = draftedMessage
      conversationState.value = 'SHOW_DRAFT'
      addMessage('ai', 'Here is the draft. You can edit it below before sending.')
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
    } catch (error) {
      addMessage(
        'ai',
        "I'm sorry, I encountered an error. Could you try describing your project again? ",
      )
    } finally {
      isLoading.value = false
    }
  }
}

const sendFinalMessage = (): void => {
  // In a real application, you would handle the form submission here.
  console.log('Final message to send:', finalMessage.value)
  conversationState.value = 'SENT'
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
        <div v-if="conversationState !== 'SENT'">
          <div ref="chatWindow" class="h-80 overflow-y-auto pr-2 space-y-4 mb-4">
            <div
              v-for="(msg, index) in chatHistory"
              :key="index"
              class="flex"
              :class="msg.sender === 'user' ? 'justify-end' : 'justify-start'"
            >
              <div
                class="p-3 rounded-lg max-w-xs"
                :class="msg.sender === 'user' ? 'bg-primary text-white' : 'bg-secondary text-white'"
              >
                {{ msg.text }}
              </div>
            </div>
          </div>

          <div v-if="conversationState === 'ASK_DETAILS'">
            <div class="flex gap-4">
              <input
                v-model="userInput"
                @keyup.enter="handleUserInput"
                :disabled="isLoading"
                type="text"
                placeholder="Type your response..."
                class="flex-grow bg-white/20 border-transparent rounded-md py-2 px-3 text-white focus:ring-accent focus:border-accent disabled:opacity-50"
              />
              <button
                @click="handleUserInput"
                :disabled="isLoading"
                class="bg-accent text-white font-bold py-2 px-5 rounded-lg hover:bg-opacity-90 transition-all disabled:opacity-50"
              >
                Send
              </button>
            </div>
          </div>

          <div v-if="conversationState === 'SHOW_DRAFT'">
            <label for="final-message" class="block text-sm font-medium text-gray-300 mb-2"
              >Here is the drafted message. Feel free to edit it before sending.</label
            >
            <textarea
              v-model="finalMessage"
              id="final-message"
              rows="6"
              class="w-full bg-white/20 border-transparent rounded-md py-2 px-3 text-white focus:ring-accent focus:border-accent"
            ></textarea>
            <button
              @click="sendFinalMessage"
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
