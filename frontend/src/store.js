import { writable } from 'svelte/store'

export const store = writable({
  error: [],
  images: [],
  imagesToUpload: [],
  modal: {
    isActive: false,
    image: ''
  },
  upload: false
});
