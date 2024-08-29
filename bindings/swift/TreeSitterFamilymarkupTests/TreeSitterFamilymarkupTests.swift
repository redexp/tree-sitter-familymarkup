import XCTest
import SwiftTreeSitter
import TreeSitterFamilymarkup

final class TreeSitterFamilymarkupTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_familymarkup())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Familymarkup grammar")
    }
}
