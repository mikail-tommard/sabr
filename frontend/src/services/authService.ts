import { fakeDelay } from './fakeApi'
import { users } from '@/mocks/data'
export const authService = {
  getUsers() { return fakeDelay(users, 150) }
}
