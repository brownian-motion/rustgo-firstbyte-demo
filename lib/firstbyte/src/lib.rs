use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn first_byte(cptr: *const libc::c_char) -> libc::c_char {
    let cstr = unsafe { CStr::from_ptr(cptr) };
    if cstr.is_empty() {
        return 0;
    }
    cstr.to_bytes()[0] as libc::c_char
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn first_byte_empty() {
        const EMPTY_STRING_WITH_NULL: &[u8] = b"\0";
        assert_eq!(
            0 as libc::c_char,
            first_byte(EMPTY_STRING_WITH_NULL.as_ptr().cast())
        )
    }

    #[test]
    fn first_byte_english() {
        const SINGLE_NAME_WITH_NULL: &[u8] = b"Jeffrey\0";
        assert_eq!(
            'J' as libc::c_char,
            first_byte(SINGLE_NAME_WITH_NULL.as_ptr().cast())
        )
    }
}
