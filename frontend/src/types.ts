export type ID = string
// Роли
// В реальном продукте роли Mentor/Admin обычно назначаются после верификации (или админом),
// поэтому на регистрации мы не спрашиваем роль (по умолчанию Student).
export type UserRole = 'Student'|'Mentor'|'Alumni'|'Admin'
export type ContentKind = 'question'|'post'

export interface User {
  id:ID; name:string; username:string; avatar:string; campusId:ID; role:UserRole; rating:number; badges:string[]
  isBanned?: boolean
}
export interface Campus { id:ID; name:string; city:string; membersCount:number; postsCount:number; questionsCount:number; description:string }
export interface Tag { id:ID; label:string; usageCount:number }
export type ContentStatus = 'active'|'hidden'|'deleted'

export interface Question {
  id:ID; type:'question'; title:string; summary:string; description:string; code?:string; tried?:string;
  authorId:ID; campusId:ID; tagIds:ID[]; createdAt:string;
  upvotes:number; answersCount:number; commentsCount:number;
  solved:boolean; bestAnswerId?:ID
  status?: ContentStatus
  pinned?: boolean
}
export interface Answer { id:ID; questionId:ID; authorId:ID; content:string; createdAt:string; upvotes:number; isBest:boolean }
// images: локальные URL (ObjectURL/Base64). В реальном backend это будут ссылки/идентификаторы файлов.
export interface Post {
  id:ID; type:'post'; title:string; summary:string; content:string;
  images?: string[];
  authorId:ID; campusId:ID; tagIds:ID[]; createdAt:string;
  likes:number; commentsCount:number;
  status?: ContentStatus
  featured?: boolean
}

// Админ: отчёты/жалобы на контент (мок)
export interface Report {
  id: ID
  createdAt: string
  reporterId: ID
  reason: 'spam'|'abuse'|'copyright'|'other'
  targetType: 'question'|'post'|'comment'|'user'
  targetId: ID
  note?: string
  status: 'open'|'in_review'|'resolved'
}
export interface Comment { id:ID; parentType:'question'|'answer'|'post'; parentId:ID; authorId:ID; content:string; createdAt:string }
export interface NotificationItem { id:ID; userId:ID; text:string; isRead:boolean; href:string }
export type FeedItem = Question | Post
