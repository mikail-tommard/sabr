import { fakeDelay } from './fakeApi'
import { answers, campuses, comments, posts, questions, tags, users } from '@/mocks/data'
// TODO: заменить на реальные API запросы (fetch/axios) при подключении backend
export const contentService = {
  getBootstrapData() { return fakeDelay({ users, campuses, tags, questions, answers, posts, comments }, 250) }
}
