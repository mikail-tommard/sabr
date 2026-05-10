import { fakeDelay } from './fakeApi'
import { reports } from '@/mocks/data'

// TODO: заменить на реальные API запросы для админки:
// GET /admin/reports, PATCH /admin/reports/:id, PATCH /admin/users/:id, ...
export const adminService = {
  getReports() {
    return fakeDelay(reports, 220)
  },
}
