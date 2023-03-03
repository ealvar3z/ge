/*	$go editor: queue.go,v 0.1 2023/03/03 05:57:15 ealvar3z Exp $ */
/*	$OpenBSD: queue.h,v 1.46 2020/12/30 13:33:12 millert Exp $	*/
/*	$NetBSD: queue.h,v 1.11 1996/05/16 05:17:14 mycroft Exp $	*/

/*
 * Copyright (c) 1991, 1993
 *	The Regents of the University of California.  All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)queue.h	8.5 (Berkeley) 8/20/94
 */

/*
 * This file ports over one data structures: tail queues.
 *
 * A tail queue is headed by a pair of pointers, one to the head of the
 * list and the other to the tail of the list. The elements are doubly
 * linked so that an arbitrary element can be removed without a need to
 * traverse the list. New elements can be added to the list before or
 * after an existing element, at the head of the list, or at the end of
 * the list. A tail queue may be traversed in either direction.
 *
 * For details on the use of these macros, see the queue(3) manual page.
 */

package main

import "unsafe"

// Tail queue definitions
type tailqEntry[T any] struct {
	next *T
	prev **T
}

type tailqHead[T any] struct {
	first *T
	last  **T
}

func TAILQ_HEAD[T any](name string, t *tailqHead[T]) *tailqHead[T] {
	return &tailqHead[T]{}
}

func TAILQ_HEAD_INITIALIZER[T any](head *tailqHead[T]) {
	head.first = nil
	head.last = &head.first
}

func TAILQ_ENTRY[T any](t *tailqEntry[T]) *tailqEntry[T] {
	return &tailqEntry[T]{}
}

// Tail queue access methods.
func (head *tailqHead[T]) TAILQ_FIRST() *T {
	return head.first
}

func (head *tailqHead[T]) TAILQ_END() *T {
	return nil
}

func (elm *tailqEntry[T]) TAILQ_NEXT() *T {
	return elm.next
}

func (head *tailqHead[T]) TAILQ_LAST(headname string) *T {
	return (*T)(unsafe.Pointer(
		uintptr(unsafe.Pointer(head.last)) +
			unsafe.Offsetof((*tailqHead[T])(nil).last)))
}

func (elm *tailqEntry[T]) TAILQ_PREV(headname string, field *tailqEntry[T]) *T {
	return (*T)(unsafe.Pointer(
		uintptr(unsafe.Pointer(elm.prev)) +
			unsafe.Offsetof(((*tailqHead[T])(nil)).last)))
}

func (head *tailqHead[T]) TAILQ_EMPTY() bool {
	return head.TAILQ_FIRST() == head.TAILQ_END()
}

func TAILQ_FOREACH[T any](_var *T,
	head *tailqHead[T],
	field *tailqEntry[T],
	body func(*T)) {
	for _var := head.TAILQ_FIRST(); _var != head.TAILQ_END(); _var = field.TAILQ_NEXT() {
		body(_var)
	}
}

func TAILQ_FOREACH_SAFE[T any](_var *T,
	head *tailqHead[T],
	field *tailqEntry[T],
	tvar **T, body func(*T)) {
	for _var := head.TAILQ_FIRST(); _var != head.TAILQ_END(); _var = *tvar {
		*tvar = field.TAILQ_NEXT()
		body(_var)
	}
}

func TAILQ_FOREACH_REVERSE[T any](_var *T,
	head *tailqHead[T],
	headname string,
	field *tailqEntry[T],
	body func(*T)) {
	for _var := head.TAILQ_LAST(headname); _var != head.TAILQ_END(); _var = field.TAILQ_PREV(headname, field) {
		body(_var)
	}
}

func TAILQ_FOREACH_REVERSE_SAFE[T any](_var *T,
	head *tailqHead[T],
	headname string,
	field *tailqEntry[T],
	tvar **T, body func(*T)) {
	for _var := head.TAILQ_LAST(headname); _var != head.TAILQ_END(); _var = *tvar {
		*tvar = field.TAILQ_PREV(headname, field)
		body(_var)
	}
}
