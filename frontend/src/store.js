import { writable } from 'svelte/store'

export const store = writable({
  error: [],
  images: [],
  imagesToUpload: [],
  upload: false
});
